package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/db"
	"time"
)

type Dao struct {
	tableName string
}

func (n *Dao) FindOne(query bson.M, obj interface{}) (interface{}, error) {
	err := n.Collection().FindOne(context.TODO(), query).Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, err
}

func (n *Dao) Aggregate(query []bson.M, obj interface{}) (interface{}, error) {
	cursor, err := n.Collection().Aggregate(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer db.CloseCursor(cursor)
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(obj)
		if err != nil {
			return nil, err
		}
	}
	return obj, err
}

func NewDao(table string) IDao {
	return &Dao{tableName: table}
}

func (n *Dao) Collection() *mongo.Collection {
	return db.Collection(n.tableName)
}

func (n *Dao) Add(param bson.M) (interface{}, error) {
	param["createTime"] = time.Now()
	param["updateTime"] = time.Now()
	return n.Collection().InsertOne(context.TODO(), param)
}

func (n *Dao) Delete(param bson.M) (interface{}, error) {
	return n.Collection().DeleteOne(context.TODO(), param)
}

func (n *Dao) Update(query bson.M, param bson.M) (interface{}, error) {
	param["updateTime"] = time.Now()
	return n.Collection().UpdateOne(context.TODO(), query, param)
}

func (n *Dao) UpdateWithOption(query bson.M, param bson.M, option *options.UpdateOptions) (interface{}, error) {
	param["updateTime"] = time.Now()
	return n.Collection().UpdateOne(context.TODO(), query, bson.M{"$set":param}, option)
}
func (n *Dao) List(query []bson.M, obj interface{}) ([]interface{}, int, error) {
	// todo
	collection := n.Collection()
	cursor, err := collection.Aggregate(context.TODO(), query)
	defer db.CloseCursor(cursor)
	if err != nil {
		return nil, 0, err
	}
	total, err := db.TotalByAll(collection, query)
	if err != nil {
		return nil, 0, err
	}
	var res []interface{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(obj)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, obj)
	}
	return res, total, nil
}
