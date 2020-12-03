package impl

import (
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/service"
)

type NodeService struct {
	nodeDao dao.INodeDao
}

func (n *NodeService) UpdateNodeInfo(heartFrom form.HeartBeatFrom) (interface{}, error) {
	panic("implement me")
}

func NewNodeService() service.INodeService{
	return &NodeService{nodeDao:daoimpl.NewNodeDao()}
}