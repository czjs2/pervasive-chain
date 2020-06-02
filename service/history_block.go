package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
)

type HistoryBlockService struct {
	dao dao.IDao
}

func (h *HistoryBlockService) LatestBlock() (interface{}, error) {
	panic("implement me")
}

func (h *HistoryBlockService) UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error) {
	param := bson.M{
		"type":     blockForm.Type,
		"number":   blockForm.Number,
		"id":       blockForm.Id,
		"height":   blockForm.Height,
		"father":   blockForm.Father,
		"hash":     blockForm.Hash,
		"vrf":      blockForm.Vrf,
		"time":     blockForm.Time,
		"interval": blockForm.Interval,
		"trans":    blockForm.Trans,
		"size":     blockForm.Size,
		"detail":   blockForm.Detail,
	}
	update := options.Update()
	update.SetUpsert(true)
	return h.dao.UpdateWithOption(bson.M{"height": blockForm.Height}, param, update)
}

func NewHistoryBlockService() IBlockService {
	return &HistoryBlockService{dao.NewDao(db.HistoryChainInfoTable)}
}
