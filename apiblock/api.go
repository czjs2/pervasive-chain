package apiblock

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/form"
	"pervasive-chain/service"
	"pervasive-chain/utils"
	"pervasive-chain/websocket"
)

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



func ChainInfoHandler(c *websocket.Client) {

}




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

