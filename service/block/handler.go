package block

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
	"time"
)

type BlockHandler struct {
	blockDao       dao.IBlockDao
	latestBlockDao dao.ILatestBlock
}

func (b *BlockHandler) WsChainInfoHandler(c *ws.WsContext) {
	fmt.Printf("websocket chain info  %v \n", time.Now())

}

func (b *BlockHandler) UpdateBlock(c *gin.Context) {
	var blockFrom ReportBlockForm
	utils.MustParams(c, &blockFrom)
	params, err := getBlockParams(blockFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	transGroup, trans := getShardTransParam(blockFrom)
	_, err = b.blockDao.Insert(params, transGroup, trans)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	utils.SuccessResponse(c, nil)
}

func (b *BlockHandler) LatestBlockInfo() {

}

func NewBlockHandler() *BlockHandler {
	return &BlockHandler{
		blockDao:       daoimpl.NewBlockDao(),
		latestBlockDao: daoimpl.NewLatestBlockDao(),
	}
}
