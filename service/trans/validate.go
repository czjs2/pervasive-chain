package trans

import (
	"pervasive-chain/service"
	"pervasive-chain/ws"
)

func TransGroupValidate(c *ws.WsContext) (service.IFormValidateInterface, error) {
	var transGroup TransGroupFrom
	return &transGroup, c.BindJSON(&transGroup)
}

func TransValidate(c *ws.WsContext) (service.IFormValidateInterface, error) {
	var trans TransFrom
	return &trans, c.BindJSON(&trans)
}
