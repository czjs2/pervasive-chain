package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDao interface {
	Add(param bson.M) (interface{}, error)
	Delete(param bson.M) (interface{}, error)
	Update(query, param bson.M) (interface{}, error)
	UpdateWithOption(query, param bson.M, update *options.UpdateOptions) (interface{}, error)
	List(query []bson.M) ([]*interface{}, error)
}
