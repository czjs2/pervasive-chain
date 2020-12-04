package ws

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	cmd "pervasive-chain/model/cmd"
	"pervasive-chain/statecode"
	"pervasive-chain/utils"
	"reflect"
	"sync"
	"time"
)

type Client struct {
	ID       string
	Socket   *websocket.Conn
	Send     chan []byte
	ClientIp string
	Dispatch *WsDispatch
	sync.Mutex
}

type WsContext struct {
	MsgId  string
	Body   string
	Uri    string
	Client *Client
}

func (c *WsContext) BindJSON(obj interface{}) error {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return fmt.Errorf("obj is mutst point")
	}
	err := json.Unmarshal([]byte(c.Body), obj)
	if err != nil {
		return fmt.Errorf("ws jsong parse error %v \n", err)
	}
	return err
}

func (c *WsContext) Json(code int, obj interface{}) {
	bytes, err := json.Marshal(cmd.NewResponseCmd(c.Uri, c.MsgId, obj))
	if err != nil {
		fmt.Printf("json marshal error %v \n", err)
		return
	}
	c.Client.Send <- bytes
}

func NewWsContext(uri, msgId, body string, c *Client) *WsContext {
	return &WsContext{Uri: uri, MsgId: msgId, Body: body, Client: c}
}

func NewClient(clientIp string, conn *websocket.Conn) *Client {
	client := &Client{
		ID:       utils.GetUUID(),
		Socket:   conn,
		Send:     make(chan []byte, 1024),
		ClientIp: clientIp,
		Dispatch: NewWsDispatch(),
	}
	WebSocketRouterV1(client)
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
		uri := utils.GetJsonValue(src, "uri")
		body := utils.GetJsonValue(src, "body")
		msgId := utils.GetJsonValue(src, "msgId")
		err = c.Dispatch.Execute(uri, NewWsContext(uri, msgId, body, c))
		if err != nil {
			bytes, err := json.Marshal(cmd.NewErrorResponse(uri, msgId, err.Error(), statecode.Fail, nil))
			if err != nil {
				// todo
				continue
			}
			c.Send <- bytes
		}
	}
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
