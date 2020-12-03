package daoimpl

import (
	"pervasive-chain/dao"
	"pervasive-chain/db"
)

type NodeDao struct {
	dao db.IDao
}

func NewNodeDao() dao.INodeDao {
	return NodeDao{dao: db.NewDaoWithTable(db.NodeTable)}
}
