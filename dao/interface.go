package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/model"
)

type IBlockDao interface {
	Insert(blockParam bson.M,transGroup,transParam []interface{}) (interface{}, error)
	Query() (interface{}, error)
}

type ITransGroupDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}

type ITransDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}


type INodeBandDao interface {

}

type ITotalBandDao interface {

}

type ILatestBlock interface {
	LatestBlockList()(interface{},error)
	UpdateBlock(chainId string,param bson.M)(interface{},error)
}


type INodeDao interface {
	FindOne(nodeId string)(*model.Node,error)

	Insert(chainType,chainKey,nodeId,latestTime string)(interface{},error)
	UpdateLatestTime(nodeId,latestTime string)(interface{},error)
}
