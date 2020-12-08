package node

import (
	"pervasive-chain/service"
	"pervasive-chain/utils"
)

func HeartBeatValidate(request string) (service.IFormValidateInterface, error) {
	obj := HeartBeatFrom{}
	err := utils.Unmarshal(request, &obj)
	return &obj, err
}

//func GenCmdValidate(request string) (service.IFormValidateInterface, error) {
//	obj := GenCmdFrom{}
//	err := utils.Unmarshal(request, &obj)
//	return &obj, err
//}
