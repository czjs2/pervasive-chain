package ws

import (
	"encoding/json"
	"pervasive-chain/model"
	"time"
)

// 订阅消息返回
func NewSubscribeResp(obj interface{}) ([]byte, error) {
	msg := model.Subscribe{
		Event: "all",
		Body:  obj,
		MsgId: time.Now().Unix(),
	}
	return json.Marshal(msg)
}

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
