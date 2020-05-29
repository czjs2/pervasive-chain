package httpsvr

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"pervasive-chain/ws"
)

// 前端订阅消息的中间件，接受到消息通过websocket直接发送到前端
func BroadcastMiddleWare(logPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
		values := c.Request.PostForm
		bytes, _ := json.Marshal(values)
		ws.BroadCast(bytes)
	}
}
