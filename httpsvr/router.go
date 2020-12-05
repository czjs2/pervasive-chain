package httpsvr

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pervasive-chain/service/block"
	"pervasive-chain/service/node"
	"pervasive-chain/ws"
	"time"
)

// websocket 路由信息
func RegisterWsRouter() ws.WsDispatch {
	dispatch := ws.NewWsDispatch()
	dispatch.Register("chainInfo", block.NewBlockHandler().WsChainInfoHandler)
	dispatch.Register("blockInfo", block.WsChainInfoHandler)
	dispatch.Register("ssInfo", block.WsChainInfoHandler)
	dispatch.Register("Block", block.WsChainInfoHandler)
	return dispatch
}

// todo 合并

// 验证其路由
func RegisterValidateRouter() {

	validateManager.Register(HeartPath, node.HeartBeatValidate)
}

// http 路由
func RegisterHttpRouter(router *gin.Engine) {
	group := router.Group("/api/v1.0/")
	group.GET(Ping, PingHandler)
	group.GET(WsConn, ws.WebSocketConnHandler)
	group.POST(Block, block.NewBlockHandler().UpdateBlock)
	group.POST(HeartPath, node.NewNodeService().UpdateNodeInfo)

}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"pong": time.Now()})
}
