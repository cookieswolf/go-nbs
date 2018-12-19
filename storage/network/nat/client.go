package nat

import (
	"context"
	"fmt"
	"github.com/NBSChain/go-nbs/storage/network/nbsnet"
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/utils"
	"github.com/golang/protobuf/proto"
	"net"
	"time"
)

var NoTimeOut = time.Time{}

const (
	KeepAliveTime    = time.Second * 100
	KeepAliveTimeOut = KeepAliveTime * 3
	BootStrapTimeOut = time.Second * 2
	ErrNoBeforeRetry = 3
	CmdTaskPoolSize  = 100
	CMDAnswerInvite  = 1
	CMDDigOut        = 2
	CMDDigSetup      = 3
)

type CmdProcess func(interface{}) error

type ClientCmd struct {
	CmdType int
	Params  interface{}
}

type Client struct {
	networkId  string
	CanServer  bool
	Ctx        context.Context
	closeCtx   context.CancelFunc
	errNo      int
	NatAddr    *nbsnet.NbsUdpAddr
	listenConn *net.UDPConn
	updateTime time.Time
	CmdTask    chan *ClientCmd
}

func NewNatClient(networkId string, canServer chan bool) (*Client, error) {

	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{
		Ctx:        ctx,
		closeCtx:   cancel,
		networkId:  networkId,
		updateTime: time.Now(),
		CmdTask:    make(chan *ClientCmd, CmdTaskPoolSize),
	}

	natServ, err := c.findWhoAmI(canServer)
	if err != nil {
		return nil, err
	}

	if c.CanServer {
		return c, nil
	}

	go c.keepAlive(natServ)
	go c.readCmd()

	return c, nil

}

func (c *Client) findWhoAmI(canSever chan bool) (*net.UDPAddr, error) {

	for _, serverIp := range utils.GetConfig().NatServerIP {

		natServerAddr := &net.UDPAddr{
			IP:   net.ParseIP(serverIp),
			Port: utils.GetConfig().NatServerPort,
		}
		l := &net.UDPAddr{
			Port: utils.GetConfig().NatClientPort,
		}
		conn, err := net.DialUDP("udp4", l, natServerAddr)
		if err != nil {
			logger.Warning("this nat server is done:->", serverIp, err)
			continue
		}

		err = conn.SetDeadline(time.Now().Add(BootStrapTimeOut))
		if err != nil {
			logger.Warning("set dead line err:->", serverIp, err)
			conn.Close()
			continue
		}

		localAddr := conn.LocalAddr().String()
		host, port, err := nbsnet.SplitHostPort(localAddr)
		request := &net_pb.NatMsg{
			Typ: nbsnet.NatBootReg,
			Seq: time.Now().Unix(),
			BootReg: &net_pb.BootReg{
				NodeId:      c.networkId,
				PrivateIp:   host,
				PrivatePort: port,
			},
		}

		requestData, _ := proto.Marshal(request)
		if no, err := conn.Write(requestData); err != nil || no == 0 {
			logger.Error("failed to send nat request to selfNatServer:->", err, no)
			conn.Close()
			continue
		}

		responseData := make([]byte, utils.NormalReadBuffer)
		hasRead, err := conn.Read(responseData)
		if err != nil {
			logger.Error("reading failed from nat server:->", err)
			conn.Close()
			continue
		}

		response := &net_pb.NatMsg{}
		if err := proto.Unmarshal(responseData[:hasRead], response); err != nil {
			logger.Error("unmarshal err:->", err)
			conn.Close()
			continue
		}

		ack := response.BootAnswer
		c.NatAddr = &nbsnet.NbsUdpAddr{
			NetworkId: c.networkId,
			CanServe:  nbsnet.CanServe(ack.NatType),
			NatServer: natServerAddr.String(),
			PubIp:     ack.PublicIp,
			NatIp:     ack.PublicIp,
			NatPort:   ack.PublicPort,
			PriIp:     host,
		}

		if ack.NatType == net_pb.NatType_ToBeChecked {

			select {
			case can := <-canSever:
				c.NatAddr.CanServe = can

			case <-time.After(BootStrapTimeOut / 2):
				c.NatAddr.CanServe = false
			}
		}
		c.CanServer = c.NatAddr.CanServe
		lisAddr := conn.LocalAddr().(*net.UDPAddr)
		conn.Close()

		if c.listenConn, err = net.ListenUDP("udp4", lisAddr); err != nil {
			logger.Warning(err)
			continue
		}

		logger.Info("create client success for network:->\n", c.NatAddr.String())
		return natServerAddr, nil
	}

	return nil, fmt.Errorf("failed to find who am I")
}

/************************************************************************
*
*			client side
*
*************************************************************************/
func (c *Client) keepAlive(natServer *net.UDPAddr) {
	request := &net_pb.NatMsg{
		Typ: nbsnet.NatKeepAlive,
		KeepAlive: &net_pb.KeepAlive{
			NodeId: c.networkId,
			LAddr:  c.listenConn.LocalAddr().String(),
		},
	}

	requestData, _ := proto.Marshal(request)

	for {
		select {
		case <-time.After(KeepAliveTime):

			if no, err := c.listenConn.WriteToUDP(requestData, natServer); err != nil || no == 0 {
				logger.Warning("failed to send keep alive channel message:", err)
				c.errNo++
			}
			logger.Debug("send  keep alive to nat server:->", c.listenConn.LocalAddr().String(), natServer)

			if c.errNo > ErrNoBeforeRetry {
				logger.Warning("too many times send errors")
				c.closeCtx()
				return
			}

		case <-c.Ctx.Done():
			logger.Info("exit sending thread cause's of context close")
			return
		}
	}
}

/************************************************************************
*
*			server side
*
*************************************************************************/

func (c *Client) refreshNatInfo(alive *net_pb.KeepAlive) {
	//TODO::
	c.updateTime = time.Now()

	if c.NatAddr.NatIp != alive.PubIP &&
		c.NatAddr.NatPort != alive.PubPort {

		c.NatAddr.NatIp = alive.PubIP
		c.NatAddr.NatPort = alive.PubPort
		logger.Info("node's nat info changed.", alive)
	}
}

func (c *Client) listenPing() {
	buffer := make([]byte, utils.NormalReadBuffer)

	n, peerAddr, err := c.listenConn.ReadFromUDP(buffer)

	logger.Warning(peerAddr, n, err)
}

func (c *Client) readCmd() {

	for {
		buffer := make([]byte, utils.NormalReadBuffer)

		n, peerAddr, err := c.listenConn.ReadFromUDP(buffer)
		if err != nil {
			if c.errNo++; c.errNo > ErrNoBeforeRetry {
				logger.Warning("too many reading error:->")
				c.closeCtx()
				return
			}
			logger.Warning("reading keep alive message failed:->", err)
			continue
		}

		msg := &net_pb.NatMsg{}
		if err := proto.Unmarshal(buffer[:n], msg); err != nil {
			logger.Warning("keep alive msg Unmarshal failed:->", err)
			continue
		}

		logger.Debug("KA c receive message:->", msg, peerAddr)

		switch msg.Typ {
		case nbsnet.NatKeepAlive:
			c.refreshNatInfo(msg.KeepAlive)
		case nbsnet.NatReversInvite:
			c.CmdTask <- &ClientCmd{
				CmdType: CMDAnswerInvite,
				Params:  msg.ReverseInvite,
			}
		case nbsnet.NatDigApply:
			c.CmdTask <- &ClientCmd{
				CmdType: CMDDigOut,
				Params:  msg.DigApply,
			}
		case nbsnet.NatDigConfirm:
			c.CmdTask <- &ClientCmd{
				CmdType: CMDDigSetup,
				Params:  msg.DigConfirm,
			}
		case nbsnet.NatPriDigSyn:
			c.priDigAck(msg, peerAddr)
		}

		select {
		case <-c.Ctx.Done():
			logger.Info("exit reading thread cause's of context close")
			return
		default:
		}
	}
}

func (c *Client) priDigAck(syn *net_pb.NatMsg, peerAddr *net.UDPAddr) {

	res := &net_pb.NatMsg{
		Typ: nbsnet.NatPriDigAck,
		Seq: syn.Seq + 1,
	}
	data, _ := proto.Marshal(res)
	logger.Info("Step 1-6:->answer dig in private:->", res)
	if _, err := c.listenConn.WriteToUDP(data, peerAddr); err != nil {
		logger.Warning("answer NatPriDigAck err:->", err)
	}
}
