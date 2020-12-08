package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/model"
)

type IBlockDao interface {
	Insert(blockParam,relayParam bson.M, transGroup, transParam []interface{}) (interface{}, error)
	Block(chainType,chainKey,hash string,height string) (interface{}, error)
	Query() (interface{}, error)
}

type ITransGroupDao interface {
	TransGroup(fromShard,toShard string,height int) (interface{}, error)
}

type ITransDao interface {
	Trans(hash string) (interface{}, error)
	Query() (interface{}, error)
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
	FindLatestOne() (*model.Node, error)
	Insert(chainType, chainKey, nodeId, latestTime string) (interface{}, error)
	UpdateLatestTime(nodeId, latestTime string) (interface{}, error)
}
