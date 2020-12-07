package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type LatestBlockDao struct {
	dao mongodb.IDao
}

func (l *LatestBlockDao) LatestBlockList() ([]*model.LatestBlock, error) {
	var res []*model.LatestBlock
	query := []bson.M{
		bson.M{"$match": bson.M{}},
	}
	_, err := l.dao.AggregateList(context.TODO(), query, func(ctx context.Context, cursor *mongo.Cursor) error {
		for cursor.Next(ctx) {
			block := &model.LatestBlock{}
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
	return &LatestBlockDao{dao: mongodb.NewDaoWithTable(mongodb.LatestBlock)}
}
