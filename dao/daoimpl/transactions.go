package daoimpl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"pervasive-chain/dao"
	"pervasive-chain/model"
	"pervasive-chain/mongodb"
)

type TransDao struct {
	dao mongodb.IDao
}

func (t *TransDao) Trans(hash string) (interface{}, error) {
	param := model.Param{}
	qeruy := []bson.M{
		bson.M{"$match": bson.M{"hash": hash}},
	}
	return param, t.dao.AggregateOne(context.TODO(), qeruy, &param)
}

func (t *TransDao) Query() (interface{}, error) {
	panic("implement me")
}

func NewTransDao() dao.ITransDao {
	return &TransDao{dao: mongodb.NewDaoWithTable(mongodb.TransactionsTable)}
}
