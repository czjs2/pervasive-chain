package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type param map[string]interface{}

type BaseCmd struct {
	Uri   string      `json:"uri"`
	Body  interface{} `json:"body"`
	MsgId string      `json:"msgId"`
}

type ErrorCmd struct {
	Uri   string      `json:"uri"`
	Body  interface{} `json:"body"`
	MsgId string      `json:"msgId"`
	Error Error       `json:"error"`
}
type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type Subscribe struct {
	Event string      `json:"event"`
	Body  interface{} `json:"body"`
	MsgId string      `json:"msgId"`
}

type EventCmd struct {
	Event string      `json:"event"`
	Body  interface{} `json:"body"`
	MsgId string      `json:"msgId"`
}

func NewEventResponse(url, msgId string) []byte {
	bytes, _ := json.Marshal(EventCmd{Event: url, MsgId: msgId, Body: gin.H{}})
	return bytes
}

func NewSubscribeResp(data interface{}) *Subscribe {
	return &Subscribe{
		Event: Block,
		Body:  data,
	}
}

func NewResponseCmd(uri, msgId string, body interface{}) BaseCmd {
	if body == nil {
		body = param{}
	}
	return BaseCmd{Uri: uri, Body: body, MsgId: msgId}
}

func NewErrorResponse(uri, msgId string, obj interface{}, code int) ErrorCmd {
	return ErrorCmd{
		Uri:   uri,
		MsgId: msgId,
		Body:  param{},
		Error: Error{
			Code:    code,
			Message: obj,
		},
	}
}
