package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/api"
	"pervasive-chain/ws"
)

func UseApiV1(router *gin.Engine) {
	group := router.Group("/api/v1.0/")
	group.POST("/wsConn", ws.WebSocketConnHandler)
	group.POST("/block", api.ReportBlockInfoHandler)
	group.POST("/headbeat", api.HeartBeatHandler)
	//group.POST("/flow", apiflow.NodeReportFlowHandler)
}


//func WebSocketRouterV1(c *ws.Client) {
//	c.Dispatch.Register("chainInfo", apiblock.WsChainInfoHandler)
//}
