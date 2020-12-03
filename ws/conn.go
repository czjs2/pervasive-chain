package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pervasive-chain/utils"
)

func WebSocketConnHandler(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{ID: utils.GetUUID(), Socket: conn, Send: make(chan []byte), ClientIp: c.ClientIP(), Dispatch: NewDisPatch()}
	Manager.Register <- client
	go client.Read()
	go client.Write()
}
