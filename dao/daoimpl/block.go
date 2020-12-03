package daoimpl

import (
	"pervasive-chain/dao"
	"pervasive-chain/db"
)

type BlockDao struct {
	dao db.IDao
}

func (b *BlockDao) Insert() (interface{}, error) {
	panic("implement me")
}

func (b *BlockDao) Query() (interface{}, error) {
	panic("implement me")
}

func NewBlockDao() dao.IBlockDao {
	return &BlockDao{dao: db.NewDaoWithTable(db.BlocksTable)}
}
