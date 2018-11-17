//+build !windows

package nat

import (
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/utils"
	"github.com/golang/protobuf/proto"
	"time"
)

func (tunnel *KATunnel) readRegResponse() error {

	responseData := make([]byte, NetIoBufferSize)
	hasRead, err := tunnel.kaConn.Read(responseData)
	if err != nil {
		logger.Warning("reading failed from nat server", err)
		return err
	}

	response := &net_pb.Response{}
	if err := proto.Unmarshal(responseData[:hasRead], response); err != nil {
		logger.Warning("unmarshal err:", err)
		return err
	}

	logger.Info("response:", response)

	resValue := response.BootRegRes

	tunnel.publicIp = resValue.PublicIp
	tunnel.publicPort = resValue.PublicPort

	go tunnel.runLoop()
	go tunnel.listening()
	go tunnel.readKeepAlive()
	return nil
}

func (tunnel *KATunnel) runLoop() {

	for {
		select {
		case <-time.After(time.Second * KeepAlive):
			tunnel.sendKeepAlive()
		case <-tunnel.closed:
			logger.Info("keep alive channel is closed.")
			return
		}
	}
}

func (tunnel *KATunnel) sendKeepAlive() error {

	request := &net_pb.NatRequest{
		MsgType: net_pb.NatMsgType_KeepAlive,
		KeepAlive: &net_pb.NatKeepAlive{
			NodeId: tunnel.networkId,
		},
	}

	requestData, err := proto.Marshal(request)
	if err != nil {
		logger.Warning("failed to marshal keep alive request:", err)
		return err
	}

	logger.Debug("keep alive channel start")

	if no, err := tunnel.kaConn.Write(requestData); err != nil || no == 0 {
		logger.Warning("failed to send keep alive channel message:", err)
		return err
	}

	return nil
}

func (tunnel *KATunnel) readKeepAlive() {

	for {
		buffer := make([]byte, utils.NormalReadBuffer)

		n, err := tunnel.kaConn.Read(buffer)
		if err != nil {
			logger.Warning("reading keep alive message failed:", err)
			continue
		}

		if err := tunnel.processResponse(buffer[:n]); err != nil {
			logger.Warning("process nat response message failed")
			continue
		}
	}
}

func (tunnel *KATunnel) processResponse(buffer []byte) error {

	response := &net_pb.Response{}
	if err := proto.Unmarshal(buffer, response); err != nil {
		logger.Warning("keep alive response Unmarshal failed:", err)
		return err
	}

	switch response.MsgType {
	case net_pb.NatMsgType_KeepAlive:
		tunnel.updateTime = time.Now()

	}

	return nil
}

func (tunnel *KATunnel) listening() {

	for {
		buffer := make([]byte, utils.NormalReadBuffer)
		n, err := tunnel.kaConn.Read(buffer)
		if err != nil {
			logger.Warning("reading keep alive message failed:", err)
			continue
		}

		if err := tunnel.processResponse(buffer[:n]); err != nil {
			logger.Warning("process nat response message failed")
			continue
		}
	}

}

func (tunnel *KATunnel) restoreNatChannel() {

}
