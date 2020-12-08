package ws

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"pervasive-chain/statecode"
	"sync"
	"time"
)

type Client struct {
	ID       string
	Socket   *websocket.Conn
	Send     chan []byte
	ClientIp string
	Dispatch WsDispatch
	CanPush  bool
	sync.Mutex
}

func NewClient(clientIp string, disPatch WsDispatch, conn *websocket.Conn) *Client {
	client := &Client{
		ID:       GetUUID(),
		Socket:   conn,
		CanPush:  false,
		Send:     make(chan []byte, 1024),
		ClientIp: clientIp,
		Dispatch: disPatch,
	}
	return client
}

func (c *Client) Ping() error {
	if _, _, err := c.Socket.NextReader(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		src := string(message)
		fmt.Printf("websocket recv %v  %v  \n", src, time.Now())
		uri := GetJsonValue(src, "uri")
		body := GetJsonValue(src, "body")
		msgId := GetJsonValue(src, "msgId")
		if !c.CanPush && uri == EventUrl {
			c.setCanPush(true)
			bytes := NewEmptyResponse(uri, msgId)
			c.Send <- bytes
		} else {
			err = c.Dispatch.Execute(uri, NewWsContext(uri, msgId, body, c))
			if err != nil {
				bytes, err := json.Marshal(NewErrorResponse(uri, msgId, err.Error(), statecode.Fail))
				if err != nil {
					// todo
					continue
				}
				c.Send <- bytes
			}
		}

	}
}

func (c *Client) setCanPush(can bool) {
	c.Lock()
	c.CanPush = can
	defer c.Unlock()
}

func (c *Client) Write() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
