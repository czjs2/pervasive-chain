package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type TransDao struct {
	dao mongodb.IDao
}

func (t *TransDao) TransactionsGroup(fromShard, toShard string, height uint64) (interface{}, error) {
	var res []model.Param
	query := []bson.M{
		bson.M{"$match": bson.M{"fromShard": fromShard, "toShard": toShard, "height": height}},
		bson.M{"$project": bson.M{"_id": 0, "hash": 1, "from": 1, "to": 1, "amount": 1}},
	}
	_, err := t.dao.AggregateList(context.TODO(), query, func(ctx context.Context, cursor *mongo.Cursor) error {
		for cursor.Next(ctx) {
			param := model.Param{}
			err := cursor.Decode(&param)
			if err != nil {
				return err
			}
			res = append(res, param)
		}
		return nil
	})
	return res, err
}

func (t *TransDao) Trans(hash string) (interface{}, error) {
	param := model.Param{}
	query := []bson.M{
		bson.M{"$match": bson.M{"hash": hash}},
	}
	return param, t.dao.AggregateOne(context.TODO(), query, &param)
}

func NewTransDao() dao.ITransDao {
	return &TransDao{dao: mongodb.NewDaoWithTable(mongodb.TransactionsTable)}
}
