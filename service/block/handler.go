package block

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/log"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
	"time"
)

type BlockHandler struct {
	blockDao       dao.IBlockDao
	latestBlockDao dao.ILatestBlock
}

func (b *BlockHandler) WsBlockInfo(c *ws.WsContext) {
	var blockForm SingleBlockForm
	_ = c.BindJSON(&blockForm)
	// todo

}

func (b *BlockHandler) WsChainInfoHandler(c *ws.WsContext) {
	log.Info("websocket chain info ", time.Now())
	latestBlockList, err := b.latestBlockDao.LatestBlockList()
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, latestBlockList)
}

func (b *BlockHandler) UpdateBlock(c *gin.Context) {
	var blockFrom ReportBlockForm
	utils.MustParams(c, &blockFrom)
	params, err := getBlockParams(blockFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	latestParams, err := getLatestParams(blockFrom)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	transGroup, trans := getShardTransParam(blockFrom)
	_, err = b.blockDao.Insert(params,latestParams, transGroup, trans)
	if err != nil {
		utils.FailResponse(c)
		return
	}
	ws.BroadcastMessage(latestParams)
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
