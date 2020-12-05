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

func ChainInfoValidate(c *ws.WsContext) (service.IFormValidateInterface, error) {
	panic(c)
}
