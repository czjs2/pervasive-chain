package block

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/log"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
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

func (b *BlockHandler) UpdateBlock(c *gin.Context) {
	var blockFrom ReportBlockForm
	utils.MustParams(c, &blockFrom)
	params, err := getBlockParams(blockFrom)
	if err != nil {
		log.Error(err.Error())
		utils.FailResponse(c)
		return
	}
	latestParams, err := getLatestParams(blockFrom)
	if err != nil {
		log.Error(err.Error())
		utils.FailResponse(c)
		return
	}
	//transGroup, trans := getTransGroupParam(blockFrom)
	//_, err = b.blockDao.Insert(params, latestParams, transGroup, trans)
	//if err != nil {
	//	utils.FailResponse(c)
	//	return
	//}
	transGroup, trans := getTransGroupParamV1(blockFrom)
	_, err = b.blockDao.InsertV1(params, latestParams, transGroup, trans)
	if err != nil {
		log.Error(err.Error())
		utils.FailResponse(c)
		return
	}
	relayBlockParam := getRealBlockParam(blockFrom)

	ws.BroadcastMessage(relayBlockParam)
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
