package ws

import (
	"encoding/json"
	"pervasive-chain/model"
)


func NewRespErr(cmd model.Cmd, errInfo string) ([]byte, error) {
	msg := model.Message{
		Uri:   cmd.Uri,
		MsgId: cmd.MsgId,
		Error: &model.Error{
			Code:    -1,
			Message: errInfo,
		},
	}
	return json.Marshal(msg)
}

func NewSuccessResp(cmd model.Cmd, obj interface{}) ([]byte, error) {
	msg := model.Message{
		Uri:   cmd.Uri,
		MsgId: cmd.MsgId,
		Body:  obj,
		Error: nil,
	}
	return json.Marshal(msg)
}
