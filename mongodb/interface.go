package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDao interface {
	Collection() *mongo.Collection

	InsertOne(ctx context.Context, param bson.M) (interface{}, error)

	InsertMany(ctx context.Context, params []interface{},opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)

	DeleteOne(ctx context.Context, param bson.M) (interface{}, error)

	DeleteMany(ctx context.Context, param bson.M) (*mongo.DeleteResult, error)

	UpdateOne(ctx context.Context, query, param bson.M) (interface{}, error)

	UnSetUpdateOne(ctx context.Context, query, param bson.M) (interface{}, error)
	// 未封装 $set
	UpdateOriginalOne(ctx context.Context, query, param bson.M) (interface{}, error)

	FindOne(ctx context.Context, query bson.M, obj interface{}) (interface{}, error)

	FindAndUpdate(ctx context.Context, query bson.M, param bson.M, update *options.FindOneAndUpdateOptions, obj interface{}) (interface{}, error)

	FindAndUpdateNoSet(ctx context.Context, query bson.M, param bson.M, update *options.FindOneAndUpdateOptions, obj interface{}) (interface{}, error)

	FindAndDelete(ctx context.Context, query bson.M, obj interface{}) (interface{}, error)

	UpdateWithOption(ctx context.Context, query, param bson.M, update *options.UpdateOptions) (interface{}, error)

	UpdateMany(ctx context.Context, query, params bson.M) (*mongo.UpdateResult, error)

	CountDocuments(ctx context.Context, query bson.M) (int64, error)

	List(ctx context.Context, query []bson.M, fn func(ctx context.Context, cursor *mongo.Cursor) error) error

	AggregateList(ctx context.Context, query []bson.M, fn func(ctx context.Context, cursor *mongo.Cursor) error) (int, error)

	// todo obj 为指针
	AggregateOne(ctx context.Context, query []bson.M, obj interface{}) error
	// 事务
	UseSession(ctx context.Context, fn func(sessionContext context.Context) error) error

	BulkWrite(ctx context.Context,models []mongo.WriteModel,opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
}
