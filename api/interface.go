package api

import "pervasive-chain/form"

type IBlockService interface {
	LatestBlockInfo() (interface{}, error)
	UpdateBlock(blockFrom form.ReportBlockForm) (interface{}, int, error)
}


type INodeService interface {
	UpdateNodeInfo(heartFrom form.HeartBeatFrom) (interface{}, error)
}


type ITransService interface {
	QueryTransByGroupId() (interface{}, error)
}
