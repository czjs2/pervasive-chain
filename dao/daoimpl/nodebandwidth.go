package daoimpl

import (
	"pervasive-chain/dao"
	"pervasive-chain/mongodb"
)

type NodeBandDao struct {
	dao mongodb.IDao
}
func NewNodeBandDao() dao.INodeBandDao{
	return &NodeBandDao{dao: mongodb.NewDaoWithTable()}
}