package nat

import (
	"github.com/NBSChain/go-nbs/storage/network/pb"
	"github.com/NBSChain/go-nbs/utils"
	"github.com/gogo/protobuf/proto"
	"net"
	"strconv"
	"sync"
	"time"
)

var logger = utils.GetLogInstance()

const NetIoBufferSize = 1 << 11
const BootStrapNatServerTimeOutInSec = 3

type nbsNatManager struct {
	sync.Mutex
	natServer     *net.UDPConn
	natType       nat_pb.NatType
	publicAddress *net.UDPAddr
	privateIP     string
	networkId     string
}

//TODO::support multiple local ip address.
func NewNatManager(networkId string) Manager {

	localPeers := ExternalIP()
	if len(localPeers) == 0 {
		logger.Panic("no available network")
	}

	logger.Debug("all network interfaces:", localPeers)

	natObj := &nbsNatManager{
		privateIP: localPeers[0],
		networkId: networkId,
	}

	if !utils.GetConfig().NatServiceOff {

		natObj.startNatService()

		go natObj.natService()
	}

	return natObj
}

//TODO:: support ipv6 later.
func (nat *nbsNatManager) startNatService() {

	natServer, err := net.ListenUDP("udp4", &net.UDPAddr{
		Port: utils.GetConfig().NatServerPort,
	})

	if err != nil {
		logger.Panic("can't start nat natServer.", err)
	}

	nat.natServer = natServer
}

func (nat *nbsNatManager) natService() {

	logger.Info(">>>>>>Nat natServer start to listen......")

	for {
		peerAddr, request, err := nat.readNatRequest()
		if err != nil {
			logger.Error(err)
		}

		if err := nat.composeNatResponse(request, peerAddr); err != nil {
			logger.Error(err)
		}
	}
}

func (nat *nbsNatManager) readNatRequest() (*net.UDPAddr, *nat_pb.BootNatRegReq, error) {

	data := make([]byte, NetIoBufferSize)

	n, peerAddr, err := nat.natServer.ReadFromUDP(data)
	if err != nil {
		logger.Warning("nat natServer read udp data failed:", err)
		return nil, nil, err
	}

	request := &nat_pb.BootNatRegReq{}
	if err := proto.Unmarshal(data[:n], request); err != nil {
		logger.Warning("can't parse the nat request", err, peerAddr)
		return nil, nil, err
	}

	logger.Debug("get nat request from client:", request)

	return peerAddr, request, nil
}

func (nat *nbsNatManager) composeNatResponse(request *nat_pb.BootNatRegReq, peerAddr *net.UDPAddr) error {

	response := &nat_pb.BootNatRegRes{}
	response.PublicIp = peerAddr.IP.String()
	response.PublicPort = string(peerAddr.Port)
	response.Zone = peerAddr.Zone

	if peerAddr.IP.Equal(net.ParseIP(request.PrivateIp)) {

		response.NatType = nat_pb.NatType_NoNatDevice
	} else if strconv.Itoa(peerAddr.Port) == request.PrivatePort {

		response.NatType = nat_pb.NatType_ToBeChecked
		nat.checkPeersNatService(peerAddr, response)
	} else {
		response.NatType = nat_pb.NatType_BehindNat
	}

	responseData, err := proto.Marshal(response)
	if err != nil {
		logger.Warning("failed to marshal nat response data", err)
		return err
	}

	if _, err := nat.natServer.WriteToUDP(responseData, peerAddr); err != nil {
		logger.Warning("failed to send nat response", err)
		return err
	}

	return nil
}

func (nat *nbsNatManager) checkPeersNatService(peerAddr *net.UDPAddr, response *nat_pb.BootNatRegRes) {

	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   peerAddr.IP,
		Port: utils.GetConfig().NatServerPort,
	})

	if err != nil {
		response.NatType = nat_pb.NatType_BehindNat
		return
	}

	conn.SetDeadline(time.Now().Add(time.Second * BootStrapNatServerTimeOutInSec))

	ping := &nat_pb.NatProbe{
		Ping: nat.networkId,
	}

	data, _ := proto.Marshal(ping)

	if _, err := conn.Write(data); err != nil {
		response.NatType = nat_pb.NatType_BehindNat
		return
	}

	responseData := make([]byte, NetIoBufferSize)
	hasRead, _, err := conn.ReadFromUDP(responseData)

	pong := &nat_pb.NatProbe{}
	if err := proto.Unmarshal(responseData[:hasRead], pong); err != nil {
		response.NatType = nat_pb.NatType_BehindNat
		return
	}

	if pong.Ping != ping.Ping {
		response.NatType = nat_pb.NatType_BehindNat
		return
	}

	response.NatType = nat_pb.NatType_CanBeNatServer
}

func (nat *nbsNatManager) connectToNatServer(serverIP string) (*net.UDPConn, error) {

	config := utils.GetConfig()
	natServerAddr := &net.UDPAddr{
		IP:   net.ParseIP(serverIP),
		Port: config.NatServerPort,
	}

	conn, err := net.DialUDP("udp", nil, natServerAddr)
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(BootStrapNatServerTimeOutInSec * time.Second))

	return conn, nil
}

func (nat *nbsNatManager) sendNatRequest(connection *net.UDPConn) error {

	localAddr := connection.LocalAddr().String()

	host, port, err := net.SplitHostPort(localAddr)
	nat.privateIP = host
	request := &nat_pb.BootNatRegReq{
		NodeId:      nat.networkId,
		PrivateIp:   host,
		PrivatePort: port,
	}

	requestData, err := proto.Marshal(request)
	if err != nil {
		logger.Error("failed to marshal nat request", err)
		return err
	}

	if no, err := connection.Write(requestData); err != nil || no == 0 {
		logger.Error("failed to send nat request to natServer ", err, no)
		return err
	}

	return nil
}

func (nat *nbsNatManager) parseNatResponse(connection *net.UDPConn) (*nat_pb.BootNatRegRes, error) {

	responseData := make([]byte, NetIoBufferSize)
	hasRead, _, err := connection.ReadFromUDP(responseData)
	if err != nil {
		logger.Error("failed to read nat response from natServer", err)
		return nil, err
	}

	response := &nat_pb.BootNatRegRes{}
	if err := proto.Unmarshal(responseData[:hasRead], response); err != nil {
		logger.Error("failed to unmarshal nat response data", err)
		return nil, err
	}

	logger.Debug("get response data from nat natServer:", response)

	port, _ := strconv.Atoi(response.PublicPort)
	nat.natType = response.NatType

	if response.NatType == nat_pb.NatType_BehindNat {
		nat.publicAddress = &net.UDPAddr{
			IP:   net.ParseIP(response.PublicIp),
			Port: port,
			Zone: response.Zone,
		}
	} else {
		nat.publicAddress = nil
	}

	return response, nil
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

//TODO:: set multiple servers to make it stronger.
func (nat *nbsNatManager) FetchNatInfo() error {

	config := utils.GetConfig()

	for _, serverIP := range config.NatServerIP {

		//TIPS:: no need to bind local host and local port right now
		connection, err := nat.connectToNatServer(serverIP)
		if err != nil {
			logger.Error("can't know who am I", err)
			goto CloseConn
		}

		if err := nat.sendNatRequest(connection); err != nil {
			goto CloseConn
		}

		_, err = nat.parseNatResponse(connection)
		if err == nil {
			break
		}

	CloseConn:
		connection.Close()
	}

	return nil
}