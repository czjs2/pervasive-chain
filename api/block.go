package api

import (
	"github.com/gin-gonic/gin"

	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/utils"
)
// 上报区块信息
func ReportBlockInfoHandler(c *gin.Context) {
	var reportBlockFrom form.ReportBlockForm
	err := c.BindJSON(&reportBlockFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	blockService := NewBlockService()
	_, code, err := blockService.UpdateBlock(reportBlockFrom)
	if err != nil {
		utils.ResponseWithCode(c, code)
		return
	}
	utils.SuccessResponse(c, nil)
}



type BlockService struct {
	blockDao       dao.IBlockDao
	latestBlockDao dao.ILatestBlock
}

func (b *BlockService) UpdateBlock(blockFrom form.ReportBlockForm) (interface{}, int, error) {
	panic(b)
}

func (b *BlockService) LatestBlockInfo() (interface{}, error) {
	return b.latestBlockDao.LatestBlockList()
}

func NewBlockService() IBlockService {
	return &BlockService{
		blockDao:       daoimpl.NewBlockDao(),
		latestBlockDao: daoimpl.NewLatestBlockDao(),
	}
}
