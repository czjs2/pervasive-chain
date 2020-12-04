package cmd

type BaseCmd struct {
	Uri   string      `json:"uri"`
	Body  interface{} `json:"body"`
	MsgId string      `json:"msgId"`
	Error ErrorMsg    `json:"error"`
}

func NewResponseCmd(uri, msgId string, obj interface{}) BaseCmd {
	return BaseCmd{
		Uri:   uri,
		MsgId: msgId,
		Body:  obj,
	}
}

func NewErrorResponse(uri, msgId, message string, code int, body interface{}) BaseCmd {
	return BaseCmd{
		Uri:   uri,
		Body:  body,
		MsgId: msgId,
		Error: ErrorMsg{
			Code:    code,
			Message: message,
		},
	}
}

type ErrorMsg struct {
	Code    int
	Message string
}


type ChainInfoCmd struct {
	Type   string `json:"type"`
	Number string `json:"number"`
	Hash   string `json:"hash"`
}
