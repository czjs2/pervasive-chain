package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDao interface {
	Add(param bson.M) (interface{}, error)

	Delete(param bson.M) (interface{}, error)

	Update(query, param bson.M) (interface{}, error)

	FindOne(query bson.M,obj interface{})(interface{},error)

	UpdateWithOption(query, param bson.M, update *options.UpdateOptions) (interface{}, error)

	// todo obj 为指针
	List(query []bson.M, obj interface{}) ([]interface{}, int, error)

	// todo obj 为指针
	Aggregate(query []bson.M,obj interface{}) (interface{}, error)

}
