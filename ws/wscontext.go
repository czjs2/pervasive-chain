package ws

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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
	bytes, err := json.Marshal(NewResponseCmd(c.Uri, c.MsgId, obj))
	if err != nil {
		fmt.Printf("json marshal error %v \n", err)
		return
	}
	c.Client.Send <- bytes
}

func NewWsContext(uri, msgId, body string, c *Client) *WsContext {
	return &WsContext{Uri: uri, MsgId: msgId, Body: body, Client: c}
}

