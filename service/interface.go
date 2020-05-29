package service

import "pervasive-chain/form"

type INodeService interface {
	UpdateNodeInfo(nodeForm form.HeartBeatFrom) (interface{}, error)
}

type IBlockService interface {
	UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error)
}

type IFlowService interface {
	UpdateFlowInfo(flowForm form.ReportFlowForm) (interface{}, error)
}
