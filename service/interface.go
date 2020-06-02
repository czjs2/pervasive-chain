package service

import (
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type IStatisticsService interface {
	// 统计流量
	CountFlow() (interface{}, error)
	// 统计交易数为tps
	CountChain() (interface{}, error)
}

//------------
type INodeService interface {
	FindAndUpdate(nodeForm form.HeartBeatFrom) (*model.Node, error)
	// 链列表
	ChainList() (interface{}, int, error)
	// 在线节点列表
	OnLineList() (interface{}, int, error)
	// 批量更新
	UpdateOnLineNodeCmd(cmd model.PyCmd)(interface{},error)
}

//--------------
type IBlockService interface {
	UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error)
	// 最新的块信息
	LatestBlock() (interface{}, error)
}

//------------
type IChainService interface {
	UpdateSharedInfo(shardForm form.ShardInfoForm) (interface{}, error)
	LatestShardInfo() (interface{}, error)
	ChainList() (interface{}, int, error)
	Chain(chainId string) (interface{}, error)
}

//---------------
type ITotalChainService interface {
	TotalFlowList() (interface{}, int, error)
	UpdateTotalChainInfo(relayNum, shardNum, nodeNum, tps int) (interface{}, error)
}

//----------
type IFlowService interface {
	UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error)
}

//---------
type ITotalFlowService interface {
	AddTotalFlow(flowForm form.TotalFlowForm) (interface{}, error)
	FlowList() (interface{}, int, error)
}
