package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TableManger struct {
	mgo    *mongo.Database
	tables []*TableInfo
}

func (tm *TableManger) CreateTable() error {
	collectionNames, err := tm.mgo.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	for i := 0; i < len(collectionNames); i++ {

	}

	for i := 0; i < len(tm.tables); i++ {

	}
}

func NewTableManager() *TableManger {
	return &TableManger{
		mgo: MongodbConn(),
	}
}

type TableInfo struct {
	TableName string
	Keys      interface{}
}
