package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/model"
)

type IBlockDao interface {
	Insert(blockParam, relayParam bson.M, transGroup, transParam []interface{}) (interface{}, error)
	InsertV1(blockParam, latestParam bson.M, transGroup, transParam interface{}) (interface{}, error)
	Block(chainType, chainKey, hash string, height uint64) (interface{}, error)
	Query() (interface{}, error)
}

type ITransGroupDao interface {
	TransGroup(fromShard, toShard string, height int) (interface{}, error)
}

type ITransDao interface {
	Trans(hash string) (interface{}, error)
	TransGroup(fromShard,toShard string,height uint64)(interface{},error)
}

type INodeBandDao interface {
}

type ITotalBandDao interface {
}

type ILatestBlock interface {
	LatestBlockList() (interface{}, error)
	UpdateBlock(chainId string, param bson.M) (interface{}, error)
}

type INodeDao interface {
	FindOne(nodeId string) (*model.Node, error)
	TotalNode(chainType string) (int, error)
	UpdateNodeCmd(chainType string,amount int) (interface{}, error)
	FindLatestOne(chainType string) (*model.Node, error)
	Insert(chainType, chainKey, nodeId string) (interface{}, error)
	UpdateLatestTime(nodeId string) (interface{}, error)
}
