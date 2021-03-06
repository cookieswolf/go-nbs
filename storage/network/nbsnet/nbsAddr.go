package nbsnet

import (
	"fmt"
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/thirdParty/gossip/pb"
	"net"
	"strconv"
)

//TODO:: refactoring this address setting.
type NbsUdpAddr struct {
	NetworkId   string
	CanServe    bool
	NetType     net_pb.NetWorkType
	NatServerIP string
	NatIp       string
	NatPort     int32
	PubIp       string
	PriIp       string
	AllPubIps   []string
}

func (addr *NbsUdpAddr) String() string {
	return fmt.Sprintf("************Nbs addr*****************\n"+
		"\tnetworkId:%20s\n"+
		"\tcanServer:%20t\n"+
		"\tNatServer:%20s\n"+
		"\tNatIp:%20s\n"+
		"\tNatPort:%20d\n "+
		"\tPubIp:%20s\n"+
		"\tPriIp:%20s\n"+
		"**************************************\n",
		addr.NetworkId, addr.CanServe, addr.NatServerIP,
		addr.NatIp, addr.NatPort, addr.PubIp, addr.PriIp)
}

func CanServe(natType net_pb.NatType) bool {

	var canService bool
	switch natType {
	case net_pb.NatType_UnknownRES:
		canService = false

	case net_pb.NatType_NoNatDevice:
		canService = true

	case net_pb.NatType_BehindNat:
		canService = false

	case net_pb.NatType_CanBeNatServer:
		canService = true

	case net_pb.NatType_ToBeChecked:
		canService = false
	}

	return canService
}

func SplitHostPort(addr string) (string, int32, error) {
	host, port, err := net.SplitHostPort(addr)
	intPort, _ := strconv.Atoi(port)

	return host, int32(intPort), err
}

func JoinHostPort(host string, port int32) string {
	portStr := strconv.Itoa(int(port))
	return net.JoinHostPort(host, portStr)
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

			//TODO:: Support ip v6 later.
			if ip = ip.To4(); ip == nil {
				continue
			}

			ips = append(ips, ip.String())
		}
	}

	return ips
}

func ConvertToGossipAddr(addr *NbsUdpAddr, nodeId string) *pb.BasicHost {
	return &pb.BasicHost{
		CanServer: addr.CanServe,
		NatServer: addr.NatServerIP,
		NatIP:     addr.NatIp,
		NatPort:   addr.NatPort,
		PriIP:     addr.PriIp,
		PubIp:     addr.PubIp,
		NetworkId: nodeId,
	}
}

func ConvertFromGossipAddr(addr *pb.BasicHost) *NbsUdpAddr {
	return &NbsUdpAddr{
		CanServe:    addr.CanServer,
		NatServerIP: addr.NatServer,
		NatIp:       addr.NatIP,
		NatPort:     addr.NatPort,
		PubIp:       addr.PubIp,
		PriIp:       addr.PriIP,
		NetworkId:   addr.NetworkId,
	}
}
