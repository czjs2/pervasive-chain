package service

import "pervasive-chain/form"

type IBlockService interface {
	BlockInfo() (interface{}, error)
	UpdateBlock(blockFrom form.ReportBlockForm) (interface{},int, error)
}

type ITransService interface {
	QueryTransByGroupId() (interface{}, error)
}
