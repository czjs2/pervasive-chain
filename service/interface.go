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

type IFlowService interface {
	UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error)
}
