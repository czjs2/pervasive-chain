package httpsvr

import (
	"github.com/gin-gonic/gin"
	"pervasive-chain/ws"
)


func UserApi(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/api/v1/")
	group.POST("conn", ws.WsConnHandler)
	return group
}
