package service

import "pervasive-chain/form"

type INodeService interface {
	UpdateNodeInfo(nodeForm form.HeartBeatFrom) (interface{}, error)
	// 链列表
	ChainList() (interface{}, int, error)
}

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

type ITotalChainService interface {
	TotalFlowList() (interface{}, int, error)
}

//----------
type IFlowService interface {
	UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error)
}

type ITotalFlowService interface {
	AddTotalFlow(flowForm form.TotalFlowForm) (interface{}, error)
	FlowList() (interface{}, int, error)
}
