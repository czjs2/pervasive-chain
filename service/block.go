package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
)

type BlockService struct {
	dao dao.IDao
}

func (b *BlockService) UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error) {
	param := bson.M{
		"type":blockForm.Type,
		"number":blockForm.Number,
		"id":blockForm.Id,
		"height":blockForm.Height,
		"father":blockForm.Father,
		"hash":blockForm.Hash,
		"vrf":blockForm.Vrf,
		"time":blockForm.Time,
		"interval":blockForm.Interval,
		"trans":blockForm.Trans,
		"size":blockForm.Size,
		"detail":blockForm.Detail,

	}
	update := options.Update()
	update.SetUpsert(true)
	return b.dao.UpdateWithOption(bson.M{"height": blockForm.Height}, bson.M{"$set": param}, update)
}

func NewBlockService() IBlockService {
	return &BlockService{dao:dao.NewDao(db.ChainInfoTable)}
}
