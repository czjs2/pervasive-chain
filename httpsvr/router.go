package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/config"
	"pervasive-chain/service/block"
	"pervasive-chain/service/node"
	"pervasive-chain/ws"
)

func RegisterWsRouter() ws.WsDispatch {
	dispatch := ws.NewWsDispatch()
	dispatch.Register(WsChainInfo, block.NewBlockHandler().WsChainInfoHandler, nil)
	//dispatch.Register(WsCmd, node.NewNodeHandler().UpdateGenerateCmd, nil)
	//dispatch.Register(WsBlockInfo, nil, nil)
	//dispatch.Register(WsSsInfo, nil, nil)
	//dispatch.Register(WsTranInfo, nil, nil)
	//dispatch.Register("blockInfo", block.WsChainInfoHandler)
	//dispatch.Register("ssInfo", block.WsChainInfoHandler)
	//dispatch.Register("Block", block.WsChainInfoHandler)
	return dispatch
}

func RegisterHttpValidateRouter() {

	validateManager.Register(HeartPath, node.HeartBeatValidate)

	validateManager.Register(Block, block.ReportBlockValidate)
}

func RegisterHttpRouter(router *gin.Engine) {
	group := router.Group(config.ApiVersion)
	group.GET(WsConn, ws.WebSocketConnHandler)
	group.POST(Block, block.NewBlockHandler().UpdateBlock)
	//group.POST(HeartPath, node.NewNodeHandler().UpdateNodeInfo)

}
