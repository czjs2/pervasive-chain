package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/api"
	"pervasive-chain/ws"
)

func UseApi(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/v1.0/")
	// socket 连接
	group.GET("conn", ws.WsConnHandler)
	// 心跳上报
	group.GET("headbeat", api.HeadBeat)
	// 区块上报
	group.GET("block", api.ReportBlock)
	// 流量上报
	group.GET("flow", api.ReportFlow)

	return group
}
