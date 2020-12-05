package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/service/block"
	"pervasive-chain/service/node"
	"pervasive-chain/ws"
)


func RegisterWsRouter() ws.WsDispatch {
	dispatch := ws.NewWsDispatch()
	dispatch.Register(WsChainInfo, block.NewBlockHandler().WsChainInfoHandler)
	//dispatch.Register("blockInfo", block.WsChainInfoHandler)
	//dispatch.Register("ssInfo", block.WsChainInfoHandler)
	//dispatch.Register("Block", block.WsChainInfoHandler)
	return dispatch
}


func RegisterValidateRouter() {

	validateManager.Register(HeartPath, node.HeartBeatValidate)

	validateManager.Register(Block, block.ReportBlockValidate)
}


func RegisterHttpRouter(router *gin.Engine) {
	group := router.Group(ApiVersion)
	group.GET(WsConn, ws.WebSocketConnHandler)
	group.POST(Block, block.NewBlockHandler().UpdateBlock)
	group.POST(HeartPath, node.NewNodeService().UpdateNodeInfo)

}

