package ws

import "encoding/json"

type BaseCmd struct {
	Uri   string  `json:"uri"`
	Body  CmdBody `json:"body"`
	MsgId string  `json:"msgId"`
}

type CmdBody struct {

}

//
type Message struct {
	Uri   string  `json:"uri"`
	Body  MsgBody `json:"body"`
	Error Error   `json:"error"`
	MsgId string  `json:"msgId"`
}

// error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type MsgBody struct {
}

func NewRespErr(cmd BaseCmd) ([]byte, error) {
	msg := Message{
		Uri:   cmd.Uri,
		MsgId: cmd.MsgId,
		Error: Error{
			Code:    -1,
			Message: "",
		},
	}
	return json.Marshal(msg)

}
