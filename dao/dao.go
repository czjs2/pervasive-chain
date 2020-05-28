package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/db"
)

type Dao struct {
	tableName string
}

func (n *Dao) List(query []bson.M) ([]*interface{}, error) {
	// todo
	panic(n)
}

func (n *Dao) Collection() *mongo.Collection {
	return db.Collection(n.tableName)
}

func (n *Dao) Add(param bson.M) (interface{}, error) {
	return n.Collection().InsertOne(context.TODO(), param)
}

func (n *Dao) Delete(param bson.M) (interface{}, error) {
	return n.Collection().DeleteOne(context.TODO(), param)
}

func (n *Dao) Update(query bson.M, param bson.M) (interface{}, error) {
	return n.Collection().UpdateOne(context.TODO(), query, param)
}

func (n *Dao) UpdateWithOption(query bson.M, param bson.M, option *options.UpdateOptions) (interface{}, error) {
	return n.Collection().UpdateOne(context.TODO(), query, param, option)
}

func NewDao(table string) IDao {
	return &Dao{tableName: table}
}
