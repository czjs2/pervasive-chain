package model

// 下发执行命令
type PyCmd struct {
	Key    string        `json:"key"`
	Params []interface{} `json:"params"`
}

type Params map[string]interface{}

type Cmd struct {
	Uri   string `json:"uri"`
	Body  PyCmd  `json:"body"`
	MsgId string `json:"msgId"`
}

type Subscribe struct {
	Event string      `json:"event"`
	Body  interface{} `json:"body"`
	MsgId int64       `json:"msgId"`
}

//
type Message struct {
	Uri   string      `json:"uri"`
	Body  interface{} `json:"body"`
	Error *Error      `json:"error"`
	MsgId string      `json:"msgId"`
}

// error
type Error struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
}
type MsgBody struct {
}
