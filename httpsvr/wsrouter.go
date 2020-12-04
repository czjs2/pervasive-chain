package httpsvr

import (
	"pervasive-chain/api/block"
	"pervasive-chain/ws"
)

// websocket 路由信息
func RegisterWsRouter() ws.WsDispatch {
	dispatch := ws.NewWsDispatch()
	dispatch.Register("/", block.WsChainInfoHandler)
	return dispatch
}
