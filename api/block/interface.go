package block

import "pervasive-chain/form"

type IBlockService interface {
	LatestBlockInfo() (interface{}, error)
	UpdateBlock(blockFrom form.ReportBlockForm) (interface{}, int, error)
}


