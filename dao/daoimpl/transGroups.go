package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type TransGroup struct {
	dao mongodb.IDao
}

func (t *TransGroup) TransGroup(fromShard, toShard string, height uint64) (interface{}, error) {
	var res []model.Param
	query := getQueryTransGroup(fromShard, toShard, height)
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
