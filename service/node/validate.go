package node

import (
	"pervasive-chain/service"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
)

func HeartBeatValidate(request string) (service.IFormValidateInterface, error) {
	obj := HeartBeatFrom{}
	err := utils.Unmarshal(request, &obj)
	return &obj, err
}

func GenCmdValidate(c *ws.WsContext)(service.IFormValidateInterface,error){
	obj:=GenCmdFrom{}
	err := c.BindJSON(&obj)
	return &obj,err
}
