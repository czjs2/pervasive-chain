package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/config"
	"pervasive-chain/service/block"
	"pervasive-chain/service/node"
	"pervasive-chain/service/trans"
	"pervasive-chain/ws"
)

func RegisterWsRouter() ws.WsDispatch {
	dispatch := ws.NewWsDispatch()
	dispatch.Register(WsChainInfo, block.NewBlockHandler().WsChainInfoHandler, nil)
	dispatch.Register(WsBlockInfo, block.NewBlockHandler().WsBlockInfo, block.SingBlockInfoValidate)
	dispatch.Register(WsCmd, node.NewNodeService().GenCmd, node.GenCmdValidate)
	dispatch.Register(WsSsInfo, trans.NewTransHandler().GransGroup, trans.TransGroupValidate)
	dispatch.Register(WsTranInfo, trans.NewTransHandler().TransInfo, trans.TransValidate)
	return dispatch
}

func RegisterHttpValidateRouter() {

	//validateManager.Register(HeartPath, node.HeartBeatValidate)

	//validateManager.Register(Block, block.ReportBlockValidate)
}

func RegisterHttpRouter(router *gin.Engine) {
	group := router.Group(config.ApiVersion)
	group.GET(WsConn, ws.WebSocketConnHandler)
	group.POST(Block, block.NewBlockHandler().UpdateBlock)
	group.POST(HeartPath, node.NewNodeService().UpdateNodeInfo)

}
