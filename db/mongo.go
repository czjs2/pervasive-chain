package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"pervasive-chain/config"
	"strings"
	"time"
)


var client *mongo.Client
var err error
var databaseName string
func InitMongo(config *config.WebConfig)error{
	var mongodbUrl string
	if config.Debug {
		mongodbUrl = config.DevMongodbUrl
	} else {
		mongodbUrl = config.MongodbUrl
	}
	databaseName = getDataBase(mongodbUrl)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrl))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func getDataBase(url string) string {
	index := strings.LastIndex(url, "/")
	return url[index+1:]
}
