package block

import (
	"pervasive-chain/service"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
)

func ReportBlockValidate(req string) (service.IFormValidateInterface, error) {
	var blockFrom ReportBlockForm
	return &blockFrom, utils.Unmarshal(req, &blockFrom)
}

func SingBlockInfoValidate(c *ws.WsContext) (service.IFormValidateInterface, error) {
	var blockFrom SingleBlockForm
	err := c.BindJSON(&blockFrom)
	return  &blockFrom,err
}
