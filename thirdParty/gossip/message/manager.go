package message

import (
	"fmt"
	"github.com/NBSChain/go-nbs/thirdParty/gossip/pb"
	"github.com/NBSChain/go-nbs/utils"
)

var logger = utils.GetLogInstance()

const (
	MsgQueuePoolSize = 1024

	MsgTypePlainTxt int32 = iota
	MsgTypeVoice
	MsgTypeImg
	MsgTypeUrl
)

type MsgEntity struct {
	MsgId   string
	From    string
	Typ     int32
	CTime   string
	PayLoad []byte
}
type MsgManager struct {
	msgQueue map[string]chan *MsgEntity
	msgCache map[string]struct{}
}

func NewMsgManager() *MsgManager {
	m := &MsgManager{
		msgQueue: make(map[string]chan *MsgEntity),
		msgCache: make(map[string]struct{}),
	}
	return m
}

func (manager MsgManager) NewSub(c string) (chan *MsgEntity, error) {
	if _, ok := manager.msgQueue[c]; ok {
		return nil, fmt.Errorf("duplicate channel subscribe to same channel:%s", c)
	}

	queue := make(chan *MsgEntity, MsgQueuePoolSize)
	manager.msgQueue[c] = queue
	return queue, nil
}

func (manager MsgManager) CancelSub(c string) {
	if queue, ok := manager.msgQueue[c]; ok {
		close(queue)
		delete(manager.msgQueue, c)
		logger.Debug("unsubscribe the channel:->", c)
	}
}
func (manager MsgManager) MsgReceiver(msg *pb.AppMsg) bool {
	if _, ok := manager.msgCache[msg.MsgId]; ok {
		return false
	}
	manager.msgCache[msg.MsgId] = struct{}{}
	queue, ok := manager.msgQueue[msg.Channel]
	if !ok {
		logger.Debug("I don't sub this channel:->", msg.Channel)
		return true
	}

	e := &MsgEntity{
		MsgId:   msg.MsgId,
		From:    msg.From,
		Typ:     msg.MsgType,
		PayLoad: msg.Payload,
		CTime:   msg.CTime,
	}

	queue <- e
	return true
}

func (manager MsgManager) runLoop() {

}
