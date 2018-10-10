package merkledag

import (
	"context"
	"errors"
	"github.com/AndreasBriese/bbloom"
	"github.com/NBSChain/go-nbs/storage/application/dataStore"
	"github.com/NBSChain/go-nbs/storage/bitswap"
	"github.com/NBSChain/go-nbs/storage/merkledag/cid"
	"github.com/NBSChain/go-nbs/storage/merkledag/ipld"
	"github.com/NBSChain/go-nbs/utils"
	"sync"
)

var instance 			*NbsDAGService
var once 			sync.Once
var parentContext 		context.Context
var logger 			= utils.GetLogInstance()
const HasBloomFilterSize	= 1 << 22
const HasBloomFilterHashes	= 7

//Can only process Block data right now.
func GetDagInstance() DAGService {
	once.Do(func() {
		parentContext 	= context.Background()

		router, err 	:= newNbsDagService()
		if err != nil {
			panic(err)
		}

		logger.Info("dag service start to run......\n")
		instance = router
	})

	return instance
}

type NbsDAGService struct {
	rehash     	bool
	checkFirst 	bool
	bloom		*bbloom.Bloom
	dataStore	dataStore.DataStore
}

func newNbsDagService() (*NbsDAGService, error) {

	bf := bbloom.New(float64(HasBloomFilterSize), float64(HasBloomFilterHashes))
	ds := dataStore.GetServiceDispatcher().ServiceByType(dataStore.ServiceTypeBlock)

	//TODO:: try to support multi protocol buffer coder.
	ipld.Register(cid.DagProtobuf, ipld.DecodeProtoBufBlock)

	return &NbsDAGService{
		checkFirst: 	true,
		rehash:     	false,
		bloom:		&bf,
		dataStore:	ds,
	}, nil
}

func (service *NbsDAGService) Has(c *cid.Cid) bool {

	key := c.Bytes()
	if ok := service.bloom.HasTS(key); ok{
		return true
	}

	keyCoded := cid.NewKeyFromBinary(key)

	if ok, _ := service.dataStore.Has(keyCoded); ok{
		service.bloom.AddTS(key)
		return true
	}

	return false
}

/*****************************************************************
*
*		DAGService interface implements.
*
*****************************************************************/
func (service *NbsDAGService) Get(cidObj *cid.Cid) (ipld.DagNode, error) {

	err := cid.ValidateCid(cidObj)
	if err != nil {
		return nil, err
	}

	key := cid.CidToDsKey(cidObj)

	data, err := service.dataStore.Get(key)

	if err == dataStore.ErrNotFound{
		return bitswap.GetSwapInstance().GetDagNode(cidObj)
	}else if err != nil{
		return nil, err
	}

	return ipld.Decode(data, cidObj)
}

func (service *NbsDAGService) GetMany([]*cid.Cid) <-chan *ipld.DagNode {
	return nil
}


func (service *NbsDAGService) Add(node ipld.DagNode) error {

	if node == nil{
		return errors.New("dag node is nil. ")
	}

	cidObj := node.Cid()
	err := cid.ValidateCid(cidObj)
	if err != nil {
		return err
	}

	if service.checkFirst && service.Has(cidObj){
		return nil
	}

	key := cid.CidToDsKey(cidObj)
	if err := service.dataStore.Put(key, node.RawData()); err != nil{
		logger.Error(err)
		return err
	}

	if err := bitswap.GetSwapInstance().SaveToNetPeer([]ipld.DagNode{node}); err != nil{
		logger.Error(err)
		return err
	}

	return nil
}


func (service *NbsDAGService) Remove(*cid.Cid) error {
	return nil
}

func (service *NbsDAGService) AddMany(nodeArr []ipld.DagNode) error {

	if len(nodeArr) == 0{
		return nil
	}

	toPut 		:= make([]ipld.DagNode, 0, len(nodeArr))
	dataBatch, err 	:= service.dataStore.Batch()
	if err != nil{
		return err
	}

	for _, node := range nodeArr{

		cidObj := node.Cid()
		if err := cid.ValidateCid(cidObj); err != nil {
			return err
		}

		if !service.checkFirst ||
			(service.checkFirst && !service.Has(cidObj)){

			toPut = append(toPut, node)

			key := cid.CidToDsKey(cidObj)
			if err := dataBatch.Put(key, node.RawData()); err != nil{
				return err
			}
		}
	}

	if err := dataBatch.Commit(); err != nil{
		return err
	}

	bitSwap := bitswap.GetSwapInstance()

	if err := bitSwap.SaveToNetPeer(toPut); err != nil{
		logger.Error(err)
	}

	for _, node := range toPut{
		service.bloom.AddTS(node.Cid().Bytes())
	}

	return nil
}

func (service *NbsDAGService) RemoveMany([]*cid.Cid) error {
	return nil
}
