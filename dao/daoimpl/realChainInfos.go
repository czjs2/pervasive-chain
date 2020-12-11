package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type LatestBlockDao struct {
	dao mongodb.IDao
}

func (l *LatestBlockDao) UpdateBlock(chainId string, param bson.M) (interface{}, error) {
	update := options.Update()
	update.SetUpsert(true)
	return l.dao.UpdateWithOption(context.TODO(), bson.M{"chainKey": chainId}, param, update)
}

func (l *LatestBlockDao) LatestBlockList() (interface{}, error) {
	var res []*model.Param
	query := []bson.M{
		bson.M{"$match": bson.M{}},
		bson.M{"$project": bson.M{"_id": 0}},
	}
	_, err := l.dao.AggregateList(context.TODO(), query, func(ctx context.Context, cursor *mongo.Cursor) error {
		for cursor.Next(ctx) {
			block := &model.Param{}
			err := cursor.Decode(block)
			if err != nil {
				return err
			}
			res = append(res, block)
		}
		return nil
	})
	return res, err
}

func NewLatestBlockDao() dao.ILatestBlock {
	return &LatestBlockDao{dao: mongodb.NewDaoWithTable(mongodb.RealChainInfosTable)}
}
