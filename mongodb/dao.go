package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pervasive-chain/utils"
	"reflect"
)

type Dao struct {
	tableName  string
	collection *mongo.Collection
}

func (n *Dao) FindAndUpdateNoSet(ctx context.Context, query bson.M, param bson.M, update *options.FindOneAndUpdateOptions, obj interface{}) (interface{}, error) {
	err := n.Collection().FindOneAndUpdate(ctx, query, param, update).Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *Dao) UnSetUpdateOne(ctx context.Context, query, param bson.M) (interface{}, error) {
	res, err := n.Collection().UpdateOne(ctx, query, bson.M{"$unset": param}, )
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (n *Dao) DeleteMany(ctx context.Context, param bson.M) (*mongo.DeleteResult, error) {
	return n.Collection().DeleteMany(ctx, param)
}

func (n *Dao) InsertMany(ctx context.Context, params []interface{}) (interface{}, error) {
	return n.Collection().InsertMany(ctx, params)
}

func (n *Dao) UpdateOriginalOne(ctx context.Context, query, param bson.M) (interface{}, error) {
	return n.Collection().UpdateOne(ctx, query, param)
}

func (n *Dao) FindAndDelete(ctx context.Context, query bson.M, obj interface{}) (interface{}, error) {
	return nil, n.Collection().FindOneAndDelete(ctx, query).Decode(obj)
}

func (n *Dao) CountDocuments(ctx context.Context, query bson.M) (int64, error) {
	return n.Collection().CountDocuments(ctx, query)
}

func (n *Dao) UseSession(ctx context.Context, fn func(ctx context.Context) error) error {
	return MongodbClient().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			errs := sessionContext.AbortTransaction(sessionContext)
			if errs != nil {
				return errs
			}
			return err
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}

func (n *Dao) UpdateMany(ctx context.Context, query, params bson.M) (*mongo.UpdateResult, error) {
	return n.Collection().UpdateMany(ctx, query, bson.M{"$set": params})
}

func (n *Dao) FindAndUpdate(ctx context.Context, query, param bson.M, update *options.FindOneAndUpdateOptions, obj interface{}) (interface{}, error) {
	err := n.Collection().FindOneAndUpdate(ctx, query, bson.M{"$set": param}, update).Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (n *Dao) FindOne(ctx context.Context, query bson.M, obj interface{}) (interface{}, error) {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("obj mutst is ptr %v", obj)
	}
	err := n.Collection().FindOne(ctx, query).Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, err
}

func NewDaoWithTable(table string) IDao {
	return &Dao{tableName: table}
}

func NewDao() IDao {
	return &Dao{}
}

func (n *Dao) Collection() *mongo.Collection {
	if n.tableName == "" {
		panic(n)
	}
	if n.collection != nil {
		return n.collection
	}
	return Collection(n.tableName)
}

func (n *Dao) InsertOne(ctx context.Context, param bson.M) (interface{}, error) {
	param["create_time"] = utils.GetMongoTime()
	param["update_time"] = utils.GetMongoTime()
	return n.Collection().InsertOne(ctx, param)
}

func (n *Dao) DeleteOne(ctx context.Context, param bson.M) (interface{}, error) {
	return n.Collection().DeleteOne(ctx, param)
}

func (n *Dao) UpdateOne(ctx context.Context, query bson.M, param bson.M) (interface{}, error) {
	param["update_time"] = utils.GetMongoTime()
	return n.Collection().UpdateOne(ctx, query, bson.M{"$set": param})
}

func (n *Dao) UpdateWithOption(ctx context.Context, query bson.M, param bson.M, option *options.UpdateOptions) (interface{}, error) {
	param["update_time"] = utils.GetMongoTime()
	return n.Collection().UpdateOne(ctx, query, bson.M{"$set": param}, option)
}

func (n *Dao) List(ctx context.Context, query []bson.M, fn func(ctx context.Context, cursor *mongo.Cursor) error) error {
	collection := n.Collection()
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := collection.Aggregate(ctx, query, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return err
	}
	return fn(ctx, cursor)
}

func (n *Dao) AggregateList(ctx context.Context, query []bson.M, fn func(context context.Context, cursor *mongo.Cursor) error) (int, error) {
	collection := n.Collection()
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := collection.Aggregate(ctx, query, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return 0, err
	}
	total, err := TotalByAll(collection, query)
	if err != nil {
		return 0, err
	}
	return int(total), fn(ctx, cursor)
}

func (n *Dao) AggregateOne(ctx context.Context, query []bson.M, obj interface{}) error {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return fmt.Errorf("obj mutst is ptr %v", obj)
	}
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := n.Collection().Aggregate(ctx, query, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return err
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewObj(obj interface{}) interface{} {
	getType := reflect.TypeOf(obj)
	getValue := reflect.ValueOf(obj)
	param := make(map[string]interface{})
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		param[field.Name] = value
	}
	return param
}
