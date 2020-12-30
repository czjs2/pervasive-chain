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
	block, err := b.blockDao.Block(blockForm.Type, blockForm.ChainKey, blockForm.Hash, blockForm.Height)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, block)

}

func (b *BlockHandler) WsChainInfoHandler(c *ws.WsContext) {
	latestBlockList, err := b.latestBlockDao.LatestBlockList()
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, gin.H{"data": latestBlockList})
}

func (b *BlockHandler) TestUpdateBlockV3(c *gin.Context) {
	start := time.Now()
	var blockFrom ReportBlockForm
	utils.MustParams(c, &blockFrom)
	params, err := getBlockParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	latestParams, err := getLatestParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	transGroup, trans := getTransGroupParam(blockFrom)
	_, err = b.blockDao.InsertV3(params, latestParams, transGroup, trans)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	relayBlockParam := getRealBlockParam(blockFrom)

	ws.BroadcastMessage(relayBlockParam)
	utils.SuccessResponse(c, nil)
	end := time.Now()
	log.Debug("spend time block:  ", end.Sub(start).Seconds(), )
}

func (b *BlockHandler) TestUpdateBlockV2(c *gin.Context) {
	start := time.Now()
	var blockFrom ReportBlockForm
	utils.MustParams(c, &blockFrom)
	params, err := getBlockParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	latestParams, err := getLatestParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	transGroup, trans := getTransGroupParam(blockFrom)
	_, err = b.blockDao.InsertV2(params, latestParams, transGroup, trans)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	relayBlockParam := getRealBlockParam(blockFrom)

	ws.BroadcastMessage(relayBlockParam)
	utils.SuccessResponse(c, nil)
	end := time.Now()
	log.Debug("spend time block:  ", end.Sub(start).Seconds(), )
}

func (b *BlockHandler) UpdateBlock(c *gin.Context) {
	var blockFrom ReportBlockForm
	err := c.BindJSON(&blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	if ok, err := blockFrom.Valid(); err != nil || !ok {
		utils.FailResponse(c, err)
		return
	}
	params, err := getBlockParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	latestParams, err := getLatestParams(blockFrom)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}
	relayBlockParam := getRealBlockParam(blockFrom)

	transGroup, trans := getTransGroupParam(blockFrom)
	_, err = b.blockDao.Insert(params, latestParams, transGroup, trans)
	if err != nil {
		utils.FailResponse(c, err.Error())
		return
	}


	//transGroup, trans := getTransGroupParamV1(blockFrom)
	//_, err = b.blockDao.InsertV1(params, latestParams, transGroup, trans)
	//if err != nil {
	//	utils.FailResponse(c,err.Error())
	//	return
	//}
	utils.SuccessResponse(c, nil)
	go ws.BroadcastMessage(relayBlockParam)


}

func (b *BlockHandler) LatestBlockInfo() {

}

func NewBlockHandler() *BlockHandler {
	return &BlockHandler{
		blockDao:       daoimpl.NewBlockDao(),
		latestBlockDao: daoimpl.NewLatestBlockDao(),
	}
}
