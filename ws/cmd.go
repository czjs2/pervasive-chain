package ws

import "encoding/json"

type Cmd struct {
	Uri   string  `json:"uri"`
	Body  CmdBody `json:"body"`
	MsgId string  `json:"msgId"`
}

type CmdBody struct {
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

func NewRespErr(cmd Cmd, errInfo string) ([]byte, error) {
	msg := Message{
		Uri:   cmd.Uri,
		MsgId: cmd.MsgId,
		Error: &Error{
			Code:    -1,
			Message: errInfo,
		},
	}
	return json.Marshal(msg)
}

func NewSuccessResp(cmd Cmd, obj interface{}) ([]byte, error) {
	msg := Message{
		Uri:   cmd.Uri,
		MsgId: cmd.MsgId,
		Body:  obj,
		Error: nil,
	}
	return json.Marshal(msg)
}
