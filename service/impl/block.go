package impl

import (
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/service"
)

type BlockService struct {
	blockDao dao.IBlockDao
}

func (b *BlockService) UpdateBlock(blockFrom form.ReportBlockForm) (interface{},int, error) {
	panic("implement me")
}

func (b *BlockService) BlockInfo() (interface{}, error) {
	panic("implement me")
}

func NewBlockService() service.IBlockService {
	return &BlockService{blockDao: daoimpl.NewBlockDao()}
}
