package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/api/block"
	"pervasive-chain/api/node"
	"pervasive-chain/ws"
)

func UseApiV1(router *gin.Engine) {
	group := router.Group("/api/v1.0/")
	group.POST("/wsConn", ws.WebSocketConnHandler)
	group.POST("/block", block.ReportBlockInfoHandler)
	group.POST("/headbeat", node.HeartBeatHandler)
	//group.POST("/flow", apiflow.NodeReportFlowHandler)
}

