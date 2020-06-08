package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/config"
	"pervasive-chain/dao"
	"pervasive-chain/db"
	"pervasive-chain/form"
	"pervasive-chain/model"
)

type BlockService struct {
	dao dao.IDao
}

func (b *BlockService) BlockList(chainType, chainId string) (interface{}, int, error) {
	query := []bson.M{
		bson.M{"$match": bson.M{"$and": []bson.M{bson.M{"type": chainType}, bson.M{"number": chainId}}}},
		bson.M{"$sort": bson.M{"height": -1}},
		bson.M{"$limit": 100},
	}
	return b.dao.List(query)
}

func (b *BlockService) ChainNodes(res []*model.Block, node *model.NodeBlock) {
	for i := 0; i < len(res); i++ {
		if res[i].Type == config.BChain {
			node.Block = res[i]
		}
	}
}

func (b *BlockService) LatestBlock() (interface{}, error) {
	query := []bson.M{
		bson.M{"$sort": bson.M{"height": -1}},
	}
	block := model.Block{}
	_, err := b.dao.Aggregate(query, &block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (b *BlockService) UpdateBlockInfo(blockForm form.ReportBlockForm) (interface{}, error) {
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
	return b.dao.UpdateWithOption(bson.M{"height": blockForm.Height, "type": blockForm.Type, "id": blockForm.Id}, param, update)
}

func NewBlockService() IBlockService {
	return &BlockService{dao: dao.NewDao(db.BlockInfoTable)}
}
