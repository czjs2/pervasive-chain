package apiblock

import (
	"pervasive-chain/dao"
	"pervasive-chain/mongodb"
)

type BlockDao struct {
	dao mongodb.IDao
}

func (b *BlockDao) Insert() (interface{}, error) {
	panic("implement me")
}

func (b *BlockDao) Query() (interface{}, error) {
	panic("implement me")
}

func NewBlockDao() dao.IBlockDao {
	return &BlockDao{dao: mongodb.NewDaoWithTable(mongodb.BlocksTable)}
}
