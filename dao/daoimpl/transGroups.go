package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type TransGroup struct {
	dao mongodb.IDao
}

func (t *TransGroup) TransGroup(fromShard, toShard string, height int) (interface{}, error) {
	param := model.Param{}
	query := []bson.M{
		bson.M{"$match": bson.M{"height": height, "fromShard": fromShard, "toShard": toShard}},
	}
	return param, t.dao.AggregateOne(context.TODO(), query, &param)
}

func (t *TransGroup) Query() (interface{}, error) {
	panic("implement me")
}

func NewTransGroup() dao.ITransGroupDao {
	return &TransGroup{
		dao: mongodb.NewDaoWithTable(mongodb.TransGroupsTable),
	}
}
