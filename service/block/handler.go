package block

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/ws"
	"time"
)

// 链整体信息
func WsChainInfoHandler(c *ws.WsContext) {
	fmt.Println("websocket recv info  %V ")
}

type BlockHandler struct {
	blockDao       dao.IBlockDao
	latestBlockDao dao.ILatestBlock
}

func (b *BlockHandler) WsChainInfoHandler(c *ws.WsContext) {
	fmt.Printf("websocket chain info  %v \n", time.Now())
}

func (b *BlockHandler) UpdateBlock(c *gin.Context) {

}

func (b *BlockHandler) LatestBlockInfo() {

}

func NewBlockHandler() *BlockHandler {
	return &BlockHandler{
		blockDao:       daoimpl.NewBlockDao(),
		latestBlockDao: daoimpl.NewLatestBlockDao(),
	}
}
