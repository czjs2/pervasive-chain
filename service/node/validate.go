package node

import (
	"fmt"
	"pervasive-chain/service"
	"pervasive-chain/utils"
)

func HeartBeatValidate(request string) (service.IFormValidateInterface, error) {
	fmt.Printf("heartBeat validate  %v  \n", request)
	obj := HeartBeatFrom{}
	err := utils.Unmarshal(request, &obj)
	return &obj, err
}
