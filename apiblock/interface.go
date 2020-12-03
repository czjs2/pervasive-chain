package apiblock

import "pervasive-chain/form"

type IBlockService interface {
	BlockInfo() (interface{}, error)
	UpdateBlock(blockFrom form.ReportBlockForm) (interface{}, int, error)
}

type IBlockDao interface {
	Insert() (interface{}, error)
	Query() (interface{}, error)
}

