package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"pervasive-chain/config"
	"pervasive-chain/model"
	"strings"
	"time"
)

var client *mongo.Client
var err error
var databaseName string
var debug bool = true

func init() {
	if debug {
		// todo just test
		mongodbUrl := "mongodb://poolwebdev:xjrw2020@139.186.84.15:27987,139.186.84.15:27988,139.186.84.15:27989/pervasivedev"
		databaseName = getDataBase(mongodbUrl)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrl))
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {

		}

	}
}

func InitMongo(config *config.WebConfig) error {
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

func Collection(collName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collName)
}

func MongoClient() *mongo.Client {
	return client
}

func CloseCursor(cursor *mongo.Cursor) {
	if cursor != nil {
		cursor.Close(context.TODO())
	}
}

func TotalByAll(collection *mongo.Collection, query []bson.M) (int, error) {
	query = append(query, bson.M{"$count": "total"})
	cursor, err := collection.Aggregate(context.TODO(), query)
	defer func() {
		if cursor != nil {
			cursor.Close(context.TODO())
		}
	}()
	if err != nil {
		return 0, err
	}
	total := model.Total{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&total)
		if err != nil {
			return 0, err
		}
	}
	return total.Total, nil
}
