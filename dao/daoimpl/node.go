package daoimpl

import (
	"pervasive-chain/dao"
	"pervasive-chain/mongodb"
)

type NodeDao struct {
	dao mongodb.IDao
}

func NewNodeDao() dao.INodeDao {
	return NodeDao{dao: mongodb.NewDaoWithTable(mongodb.NodeTable)}
}
