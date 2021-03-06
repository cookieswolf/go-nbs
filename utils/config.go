package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"sync"
)

type Configure struct {
	BaseDir            string
	LevelDBDir         string
	BlocksDir          string
	ShardFun           string
	LogFileName        string
	CmdServicePort     string
	CurrentVersion     string
	SysTimeFormat      string
	BootStrapPeers     []string
	NatServerPort      int
	HolePuncherPort    int
	NatPrivatePingPort int
	NatServerIP        []string
	GossipBootStrapIP  []string
	GossipCtrlPort     int
}

const (
	currentVersion     = "0.1.4"
	cmdServicePort     = 10001
	natServerPort      = 11001
	natChanPriPingPort = 11002
	holePuncherPort    = 11003
	gossipCtrlPort     = 12001
	NormalReadBuffer   = 1 << 11
	AdditionalCopies   = 1
)

var (
	natServerIP = []string{
		"155.138.212.125",
		"103.45.98.72",
		//"192.168.103.97",
		//"192.168.103.101",
	}

	gossipContracts = []string{
		"103.45.98.72",
		"155.138.212.125",
		//"192.168.103.97",
		//"192.168.103.101",
	}

	defaultBootstrapAddresses = []string{
		"",
	}

	config   *Configure
	onceConf sync.Once
)

//TODO:: config to local storage.
func GetConfig() *Configure {
	onceConf.Do(func() {
		config = initConfig()
	})

	return config
}

func initConfig() *Configure {

	baseDir := getBaseDir()
	if _, ok := FileExists(baseDir); ok == false {
		err := os.Mkdir(baseDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	levelDBDir := filepath.Join(baseDir, string(filepath.Separator), "dataStore")
	blockStoreDir := filepath.Join(baseDir, string(filepath.Separator), "blocks")
	logFileName := filepath.Join(baseDir, string(filepath.Separator), "nbs.log")

	return &Configure{
		BaseDir:            baseDir,
		LevelDBDir:         levelDBDir,
		BlocksDir:          blockStoreDir,
		ShardFun:           "/repo/flatfs/shard/v1/next-to-last/2",
		LogFileName:        logFileName,
		CmdServicePort:     strconv.Itoa(cmdServicePort),
		CurrentVersion:     currentVersion,
		SysTimeFormat:      "2006-01-02 15:04:05",
		BootStrapPeers:     defaultBootstrapAddresses,
		NatServerPort:      natServerPort,
		NatPrivatePingPort: natChanPriPingPort,
		HolePuncherPort:    holePuncherPort,
		NatServerIP:        natServerIP,
		GossipBootStrapIP:  gossipContracts,
		GossipCtrlPort:     gossipCtrlPort,
	}
}

func getBaseDir() string {

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	baseDir := filepath.Join(usr.HomeDir, string(filepath.Separator), ".nbs")

	return baseDir
}
