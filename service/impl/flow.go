package impl

import (
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/service"
)

type FlowService struct {
	flowDao      dao.INodeBandDao
	totalFlowDao dao.ITotalBandDao
}

func NewFlowService() service.IFlowService{
	return &FlowService{
		flowDao:daoimpl.NewBlockDao(),
	}
}
