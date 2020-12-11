package trans

import (
	"pervasive-chain/dao"
	"pervasive-chain/dao/daoimpl"
	"pervasive-chain/utils"
	"pervasive-chain/ws"
)

type TransHandler struct {
	transDao      dao.ITransDao
	transGroupDao dao.ITransGroupDao
}

func (t *TransHandler) GransGroup(c *ws.WsContext) {
	var transGroupFrom TransGroupFrom
	_ = c.BindJSON(&transGroupFrom)
	transGroup, err := t.transDao.TransGroup(transGroupFrom.FromShard, transGroupFrom.ToShard, transGroupFrom.Height)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, transGroup)
}

func (t *TransHandler) TransInfo(c *ws.WsContext) {
	var transFrom TransFrom
	_ = c.BindJSON(&transFrom)
	transGroup, err := t.transDao.Trans(transFrom.Hash)
	if err != nil {
		utils.WsFailResponse(c)
		return
	}
	utils.WsSuccessResponse(c, transGroup)
}

func NewTransHandler() *TransHandler {
	return &TransHandler{
		transDao:      daoimpl.NewTransDao(),
		transGroupDao: daoimpl.NewTransGroup(),
	}
}
