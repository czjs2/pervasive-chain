package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/api"
	"pervasive-chain/websocket"
)

func UseApiV1(router *gin.Engine) {
	group := router.Group("/api/v1.0/")
	group.POST("/wsConn", websocket.WebSocketConnHandler)
	group.POST("/block", api.ReportBlockInfoHandler)
	group.POST("/headbeat", api.HeartBeatHandler)
	group.POST("/flow", api.HeartBeatHandler)
}
