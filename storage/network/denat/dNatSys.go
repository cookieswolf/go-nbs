package denat

import (
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/utils"
	"net"
	"strconv"
	"sync"
)

type DecenterNatSys interface {
	Setup(networkId string)
	GetValidServer() string
	BroadCast(request *net_pb.DeNatSysReq) *net_pb.DeNatSysRsp
}

//decentralized nat server = dns
type ServerNode struct {
	networkID string
	hosts     []string
}

var (
	instance *ServerNode
	once     sync.Once
)

func GetDeNatSerIns() DecenterNatSys {
	once.Do(func() {
		instance = newDeNatSer()
	})

	return instance
}

func newDeNatSer() *ServerNode {

	officerServer := utils.GetConfig().NatServerIP
	server := &ServerNode{
		hosts: make([]string, len(officerServer)),
	}

	port := strconv.Itoa(utils.GetConfig().NatServerPort)

	for i, host := range officerServer {

		server.hosts[i] = net.JoinHostPort(host, port)
	}

	return server
}

func (s *ServerNode) Setup(networkId string) {
	s.networkID = networkId
}

//TODO:: use gossip protocol to manager all nat servers. we use official nat servers right now.
func (s *ServerNode) GetValidServer() string {
	return s.hosts[0] //TIPS:: simply use the first server.
}

//find client item from other nat server by peerId
func (s *ServerNode) BroadCast(request *net_pb.DeNatSysReq) *net_pb.DeNatSysRsp {
	return nil
}
