package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type TransGroup struct {
	dao mongodb.IDao
}

func (t *TransGroup) TransGroup(fromShard, toShard string, height int) (interface{}, error) {
	var res []model.Param
	query := []bson.M{
		bson.M{"$match": bson.M{"height": height, "fromShard": fromShard, "toShard": toShard}},
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

func (t *TransGroup) Query() (interface{}, error) {
	panic("implement me")
}

func NewTransGroup() dao.ITransGroupDao {
	return &TransGroup{
		dao: mongodb.NewDaoWithTable(mongodb.TransGroupsTable),
	}
}
