package node

import (
	"fmt"
	"pervasive-chain/form"
	"pervasive-chain/utils"
)

func HeartBeatValidate(request string) (form.IFormValidateInterface, error) {
	fmt.Printf("heartBeat validate  %v  \n", request)
	obj := form.HeartBeatFrom{}
	err := utils.Unmarshal(request, &obj)
	return &obj, err
}
