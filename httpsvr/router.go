package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/ws"
)


func UseApi(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/api/v1/")
	group.GET("conn", ws.WsConnHandler)
	return group
}
