package dao

import "pervasive-chain/model"

type IBlockDao interface {
	Insert() (interface{}, error)
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
	LatestBlockList()([]*model.LatestBlock,error)
}



type INodeDao interface {
	FindOne(nodeId string)(*model.Node,error)
	UpdateLatestTime(nodeId,latestTime string)(interface{},error)
}
