package nat

import (
	"github.com/NBSChain/go-nbs/storage/network/denat"
	"github.com/NBSChain/go-nbs/storage/network/nbsnet"
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/storage/network/shareport"
	"github.com/NBSChain/go-nbs/utils"
	"github.com/golang/protobuf/proto"
	"net"
	"strconv"
	"time"
)

func NewNatManager(networkId string) *Manager {

	denat.GetDNSInstance().Setup(networkId)

	localPeers := ExternalIP()
	if len(localPeers) == 0 {
		logger.Panic("no available network")
	}

	logger.Debug("all network interfaces:", localPeers)

	natObj := &Manager{
		networkId: networkId,
		canServe:  make(chan bool),
		cache:     make(map[string]*HostBehindNat),
	}

	natObj.startNatService()

	go natObj.natServiceListening()

	go natObj.cacheManager()

	return natObj
}

func (nat *Manager) SetUpNatChannel() error {

	port := strconv.Itoa(utils.GetConfig().NatChanSerPort)
	listener, err := shareport.ListenUDP("udp4", "0.0.0.0:"+port)
	if err != nil {
		logger.Warning("create share listening udp failed.")
		return err
	}

	serverIP := denat.GetDNSInstance().GetValidServer()
	client, err := shareport.DialUDP("udp4", "0.0.0.0:"+port, serverIP)
	if err != nil {
		logger.Warning("create share port dial udp connection failed.")
		return err
	}

	tunnel := &KATunnel{
		networkId:  nat.networkId,
		closed:     make(chan bool),
		serverHub:  listener,
		kaConn:     client,
		sharedAddr: client.LocalAddr().String(),
		updateTime: time.Now(),
		natTask:    make(map[string]*nbsnet.ConnTask),
	}

	go tunnel.runLoop()

	go tunnel.listening()

	go tunnel.readKeepAlive()

	nat.NatKATun = tunnel

	return nil
}

func (nat *Manager) WaitNatConfirm() chan bool {
	return nat.canServe
}

func (nat *Manager) MakeAReverseNatConn(lAddr, rAddr *nbsnet.NbsUdpAddr, connId string) (*nbsnet.ConnTask, error) {

	connTask := &nbsnet.ConnTask{
		SessionId: connId,
		Err:       make(chan error),
	}

	if lAddr.PriPort == 0 {
		lAddr.PriPort = utils.GetConfig().NatChanSerPort

		connTask.CType = nbsnet.CTypeNatReverseDirect

	} else {
		proxyConn, err := net.ListenUDP("udp4", &net.UDPAddr{
			Port: lAddr.PriPort,
			IP:   net.ParseIP(lAddr.PriIp),
		})

		if err != nil {
			return nil, err
		}

		connTask.ProxyConn = proxyConn
		connTask.CType = nbsnet.CTypeNatReverseWithProxy
	}

	if err := nat.NatKATun.invitePeer(connTask, lAddr, rAddr, connId); err != nil {
		return nil, err
	}

	return connTask, nil
}

func (nat *Manager) PunchANatHole(lAddr, rAddr *nbsnet.NbsUdpAddr, connId string) (*nbsnet.ConnTask, error) {

	task := &nbsnet.ConnTask{
		sessionId: connId,
	}

	payload := &net_pb.NatConReq{
		FromPeerId: fromId,
		ToPeerId:   toId,
		ToPort:     int32(port),
		SessionId:  sessionId,
	}
	request := &net_pb.NatRequest{
		MsgType: net_pb.NatMsgType_Connect,
		ConnReq: payload,
	}

	reqData, err := proto.Marshal(request)
	if err != nil {
		logger.Error("failed to marshal the nat connect request", err)
		task.Err = err
		return task
	}

	if no, err := tunnel.kaConn.Write(reqData); err != nil || no == 0 {
		logger.Warning("nat channel keep alive message failed", err, no)
		task.Err = err
		return task
	}

	task.ConnCh = make(chan *net.UDPConn)

	tunnel.natTask[connId] = task

	return task
}

func ExternalIP() []string {

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var ips []string
	for _, face := range interfaces {

		if face.Flags&net.FlagUp == 0 ||
			face.Flags&net.FlagLoopback != 0 {
			continue
		}

		address, err := face.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range address {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			//TODO:: Support ip v6 lter.
			if ip = ip.To4(); ip == nil {
				continue
			}

			ips = append(ips, ip.String())
		}
	}

	return ips
}
