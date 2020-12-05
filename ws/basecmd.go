package ws

type BaseCmd struct {
	Uri   string `json:"uri"`
	Body  interface{} `json:"body"`
	MsgId string `json:"msgId"`
}

type ErrorCmd struct {
	Uri   string `json:"uri"`
	Body  string `json:"body"`
	MsgId string `json:"msgId"`
	Error Error  `json:"error"`
}
type Error struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func NewResponseCmd(uri, msgId string,body interface{}) BaseCmd {
	return BaseCmd{Uri: uri, Body: body, MsgId: msgId}
}

func NewErrorResponse(uri, body, msgId string, code int) ErrorCmd {
	return ErrorCmd{
		Uri:   uri,
		Body:  body,
		MsgId: msgId,
		Error: Error{Code: code},
	}
}
