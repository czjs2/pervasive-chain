package daoimpl

import (
	"pervasive-chain/dao"
	"pervasive-chain/db"
)

type NodeBandDao struct {
	dao db.IDao
}
func NewNodeBandDao() dao.INodeBandDao{
	return &NodeBandDao{dao:db.NewDaoWithTable()}
}