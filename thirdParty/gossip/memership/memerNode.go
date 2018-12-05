package memership

import (
	"crypto/rand"
	"fmt"
	"github.com/NBSChain/go-nbs/storage/network"
	"github.com/NBSChain/go-nbs/storage/network/nbsnet"
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/thirdParty/gossip/pb"
	"github.com/NBSChain/go-nbs/utils"
	"github.com/golang/protobuf/proto"
	"math/big"
	"net"
	"time"
)

const (
	MemShipHeartBeat = time.Second * 10 //TODO::?? heart beat time interval.
	MaxInnerTaskSize = 1 << 10
	MaxForwardTimes  = 10
	DefaultSubExpire = time.Hour
	SubscribeTimeOut = time.Second * 2
	IsolatedTime     = MemShipHeartBeat * 3
)

var (
	HandlerNotFound         = fmt.Errorf("no such gossip task handler")
	PartialViewItemNotFound = fmt.Errorf("no such item in partial view")
)

type msgTask struct {
	msg  *pb.Gossip
	addr *net.UDPAddr
}

type worker func(*msgTask) error

type MemManager struct {
	nodeID      string
	updateTime  time.Time
	taskQueue   chan *msgTask
	serviceConn *nbsnet.NbsUdpConn
	inputView   map[string]*viewNode
	partialView map[string]*viewNode
	taskRouter  map[net_pb.MsgType]worker
}

var (
	logger = utils.GetLogInstance()
)

func NewMemberNode(peerId string) *MemManager {

	node := &MemManager{
		nodeID:      peerId,
		taskQueue:   make(chan *msgTask, MaxInnerTaskSize),
		inputView:   make(map[string]*viewNode),
		partialView: make(map[string]*viewNode),
		taskRouter:  make(map[net_pb.MsgType]worker),
	}

	node.taskRouter[nbsnet.GspSub] = node.firstInitSub
	node.taskRouter[nbsnet.GspVoteContact] = node.getVoteApply
	node.taskRouter[nbsnet.GspVoteResult] = node.subToContract
	node.taskRouter[nbsnet.GspHeartBeat] = node.getHeartBeat
	node.taskRouter[nbsnet.GspIntroduce] = node.getForwardSub

	return node
}

func (node *MemManager) InitNode() error {

	if err := node.initMsgService(); err != nil {
		return err
	}

	go node.receivingCmd()

	go node.RunLoop()

	if err := node.RegisterMySelf(); err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}

func (node *MemManager) initMsgService() error {

	conn, err := network.GetInstance().ListenUDP("udp4", &net.UDPAddr{
		Port: utils.GetConfig().GossipCtrlPort,
	})

	if err != nil {
		logger.Error("can't start contract service:", err)
		return err
	}

	node.serviceConn = conn

	return nil
}

func (node *MemManager) receivingCmd() {

	for {
		buffer := make([]byte, utils.NormalReadBuffer)

		n, peerAddr, err := node.serviceConn.ReceiveFromUDP(buffer)
		if err != nil {
			logger.Warning("reading contact application err:", err)
			continue
		}

		message := &pb.Gossip{}
		if err := proto.Unmarshal(buffer[:n], message); err != nil {
			logger.Warning("this is not a gossip message:->", buffer)
			continue
		}

		logger.Debug("gossip server:->", peerAddr, message)

		node.taskQueue <- &msgTask{
			msg:  message,
			addr: peerAddr,
		}
	}
}

func (node *MemManager) RunLoop() {

	for {
		select {
		case task := <-node.taskQueue:

			msgType := task.msg.MsgType
			handler, ok := node.taskRouter[msgType]
			if !ok {
				logger.Error("gossip msg handler err:->", HandlerNotFound)
			}
			if err := handler(task); err != nil {
				logger.Error("gossip run loop err:->", err)
			}

		case <-time.After(MemShipHeartBeat):
			node.sendHeartBeat()
		}
	}
}

//TODO:: make sure it's fine

func (node *MemManager) sendHeartBeat() {

	now := time.Now()
	if now.Sub(node.updateTime) > IsolatedTime {
		node.Resub()
	}

	msg := &pb.Gossip{
		MsgType: nbsnet.GspHeartBeat,
		HeartBeat: &pb.HeartBeat{
			FromID: node.nodeID,
		},
	}

	data, _ := proto.Marshal(msg)
	for _, item := range node.partialView {
		if !item.needUpdate() {
			continue
		}
		if err := item.sendData(data); err != nil {
			node.removeFromView(item, node.partialView)
		}

		if now.After(item.expiredTime) {
			node.removeFromView(item, node.partialView)
			node.unsubItem(item) //TODO::???
		}
	}
}

func (node *MemManager) getForwardSub(task *msgTask) error {
	req := task.msg.Subscribe

	prob := float64(1) / float64(1+len(node.partialView))
	random, _ := rand.Int(rand.Reader, big.NewInt(100))
	//TODO:: make sure this probability is fine.
	if random.Int64() < int64(prob*100) {
		return node.asSubAdapter(req)
	}

	if item := node.choseRandomInPartialView(); item != nil {
		return item.send(task.msg)
	}
	return node.asSubAdapter(req)
}
