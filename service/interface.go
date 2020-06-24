package service

import (
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type IStatisticsService interface {
	// 统计整个流量
	CountFlow(flowForm form.ReportFlowForm) (interface{}, error)
	// 统计某个交易的数为tps
	CountChain(chainId, chainType string) (interface{}, error)
	// 统计节点总数
	CountNode() (interface{}, error)
	// 系统整个tps
	CountTps() (interface{}, error)
	// 所有链关系
	AllChain() (interface{}, error)

	// ws chainInfo 接口
	ChainInfo() (interface{}, error)
	BlockInfo(cType, number string) (interface{}, error)
}

//------------
type INodeService interface {
	FindAndUpdate(nodeForm form.HeartBeatFrom) (*model.Node, error)
	// 链列表
	ChainList() (interface{}, int, error)
	// 在线节点列表
	OnLineList() (interface{}, int, error)
	// 批量更新
	UpdateOnLineNodeCmd(cmd model.PyCmd) (interface{}, error)

	ClearCmd() (interface{}, error)

	LatestNodeCmd() (*model.Node, error)
}

//--------------
type IBlockService interface {
	UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error)
	// 最新的块信息
	LatestBlock() (interface{}, error)
	// 获取系统链的信息
	BlockList(chainType, chainId string) (interface{}, int, error)
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
