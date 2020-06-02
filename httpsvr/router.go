package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/api"
	"pervasive-chain/ws"
)

func UseApi(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/v1.0/")
	// socket 连接
	group.GET("conn", ws.WebSocketConnHandler)
	// 心跳上报
	group.POST("headbeat", api.ReportHeadBeatHandler)
	// 区块上报
	group.POST("block", api.ReportBlockHandler)
	// 流量上报
	group.POST("flow", api.ReportFlowHandler)

	return group
}
