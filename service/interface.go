package service

import "pervasive-chain/form"

type IBlockService interface {
	BlockInfo() (interface{}, error)
	UpdateBlock(blockFrom form.ReportBlockForm) (interface{}, int, error)
}

type INodeService interface {
	UpdateNodeInfo(heartFrom form.HeartBeatFrom) (interface{}, error)
}

type IFlowService interface {
	UpdateFlow(flowFrom form.ReportFlowForm)(interface{},error)
}

type ITransService interface {
	QueryTransByGroupId() (interface{}, error)
}
