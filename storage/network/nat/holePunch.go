package nat

import (
	"fmt"
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

//conn inviter call first
func (tunnel *KATunnel) natHoleStep1InvitePeer(lAddr, rAddr *nbsnet.NbsUdpAddr, connId string) error {

	connReq := &net_pb.NatConInvite{

		FromAddr: &net_pb.NbsAddr{
			NetworkId: lAddr.NetworkId,
			CanServer: lAddr.CanServe,
			PriIp:     lAddr.PriIp,
			PriPort:   lAddr.PriPort,
			NatIP:     lAddr.NatIp,
			NatPort:   lAddr.NatPort,
		},
		ToAddr: &net_pb.NbsAddr{
			NetworkId: rAddr.NetworkId,
			CanServer: rAddr.CanServe,
			PriIp:     rAddr.PriIp,
			PriPort:   rAddr.PriPort,
			NatIP:     rAddr.NatIp,
			NatPort:   rAddr.NatPort,
		},
		SessionId: connId,
	}

	response := &net_pb.NatRequest{
		MsgType: net_pb.NatMsgType_Connect,
		ConnReq: connReq,
	}

	toItemData, err := proto.Marshal(response)
	if err != nil {
		return err
	}

	if _, err := tunnel.kaConn.Write(toItemData); err != nil {
		return err
	}

	return nil
}

//caller make a direct connection to peer's public address
func (tunnel *KATunnel) natHoleStep2Call(sessionId string, rAddr *nbsnet.NbsUdpAddr) (*net.UDPConn, error) {
	port := strconv.Itoa(int(rAddr.NatPort))
	remoteAddr := net.JoinHostPort(rAddr.NatIp, port)

	conn, err := tunnel.digIn(sessionId, remoteAddr)
	if err == nil {
		return conn, nil
	}

	conn.Close()
	remoteAddr = net.JoinHostPort(rAddr.PriIp, port)
	return tunnel.digIn(sessionId, remoteAddr)
}

//TIPS::get peer's addr info and make a connection.
func (tunnel *KATunnel) natHoleStep4Answer(response *net_pb.NatConInvite) {

	sessionId := response.SessionId

	pubAddr := &net.UDPAddr{
		IP:   net.ParseIP(response.FromAddr.NatIP),
		Port: int(response.FromAddr.NatPort),
	}
	priAddr := &net.UDPAddr{
		IP:   net.ParseIP(response.FromAddr.PriIp),
		Port: int(response.FromAddr.NatPort),
	}

	holeMsg := &net_pb.NatRequest{
		MsgType: net_pb.NatMsgType_DigOut,
		HoleMsg: &net_pb.HoleDig{
			SessionId: sessionId,
		},
	}

	holeData, _ := proto.Marshal(holeMsg)

	for i := 0; i < HolePunchingTimeOut/2; i++ {

		_, errPub := tunnel.serverHub.WriteToUDP(holeData, pubAddr)
		_, errPri := tunnel.serverHub.WriteToUDP(holeData, priAddr)

		if errPri != nil && errPub != nil {
			logger.Error("failed to dig out :", errPub, errPri)
			break
		}
		time.Sleep(time.Second)
	}
}

func (tunnel *KATunnel) digIn(sessionId string, remoteAddr string) (*net.UDPConn, error) {

	conn, err := shareport.DialUDP("udp4", tunnel.sharedAddr, remoteAddr)
	if err != nil {
		logger.Info("failed to setup hole connection by public nat ip")
		return nil, err
	}

	holeMsg := &net_pb.NatRequest{
		MsgType: net_pb.NatMsgType_DigIn,
		HoleMsg: &net_pb.HoleDig{
			SessionId: sessionId,
		},
	}

	data, _ := proto.Marshal(holeMsg)
	buffer := make([]byte, utils.NormalReadBuffer)

	for i := 0; i < HolePunchingTimeOut/2; i++ {

		if _, err := conn.Write(data); err != nil {
			logger.Debug("dig hole failed:", err)
		}
		conn.SetReadDeadline(time.Now().Add(time.Second))

		if _, err := conn.Read(buffer); err != nil {
			logger.Debug("read hole msg:", err)
			continue
		}

		return conn, nil
	}

	return nil, fmt.Errorf("time out")
}

/************************************************************************
*
*			server side
*
*************************************************************************/
//TIPS:: the server forward the connection invite to peer
func (nat *Manager) natHoleStep3ForwardInvite(request *net_pb.NatRequest, peerAddr *net.UDPAddr) error {

	req := request.ConnReq
	rawData, _ := proto.Marshal(request)

	nat.cacheLock.Lock()
	defer nat.cacheLock.Unlock()

	toItem, ok := nat.cache[req.ToAddr.NetworkId]
	if !ok {
		if err := denat.GetDeNatSerIns().ProxyConnInvite(req); err != nil {
			return err
		}
	} else {
		if _, err := nat.sysNatServer.WriteToUDP(rawData, toItem.pubAddr); err != nil {
			return err
		}
	}

	return nil
}
