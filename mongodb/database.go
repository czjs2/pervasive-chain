package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"pervasive-chain/model"
	"strings"
	"time"
)

var client *mongo.Client
var err error
var DatabaseName string

// todo
var Debug = true

func init() {
	if Debug {
		var mongodbUrl string = "mongodb://pynxtest:xjrw2020@118.24.168.230:27017,118.24.168.230:27018,118.24.168.230:27019/pynxtest"
		DatabaseName = getDataBase(mongodbUrl)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrl).SetMaxPoolSize(20))
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func MongodbInit(config *model.RuntimeConfig) error {
	var mongodbUrl string
	if config.Debug {
		mongodbUrl = config.DevMongodbUrl
	} else {
		mongodbUrl = config.MongodbUrl
	}
	Debug = config.Debug
	DatabaseName = getDataBase(mongodbUrl)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrl))
	if err!=nil{
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

// 获取列表的总数 ，手动查两次，可以使用 $facet mongodb自动查两次
func Total(collection *mongo.Collection, query bson.M) (int64, error) {
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{
		query,
		bson.M{"$group": bson.M{"_id": "", "total": bson.M{"$sum": 1}}},
	}, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return 0, err
	}
	total := TotalCount{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&total)
		if err != nil {
			return 0, err
		}
	}
	return total.Total, nil
}

func TotalManyQuery(collection *mongo.Collection, query ...bson.M) (int64, error) {
	query = append(query, bson.M{"$count": "total"})
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := collection.Aggregate(context.TODO(), query, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return 0, err
	}
	total := TotalCount{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&total)
		if err != nil {
			return 0, err
		}
	}
	return total.Total, nil
}

func TotalByAll(collection *mongo.Collection, query []bson.M) (int64, error) {
	newQuery, err := deleteSkipOrLimit(query)
	if err != nil {
		return 0, err
	}
	newQuery = append(newQuery, bson.M{"$count": "total"})
	aggregate := options.Aggregate()
	aggregate.SetAllowDiskUse(true)
	cursor, err := collection.Aggregate(context.TODO(), newQuery, aggregate)
	defer CloseCursor(cursor)
	if err != nil {
		return 0, err
	}
	total := TotalCount{}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&total)
		if err != nil {
			return 0, err
		}
	}
	return total.Total, nil
}

func deleteSkipOrLimit(query []bson.M) ([]bson.M, error) {
	var newQuery []bson.M
	for i := 0; i < len(query); i++ {
		bm := query[i]
		if _, ok := bm["$limit"]; ok {
			continue
		}
		if _, ok := bm["$skip"]; ok {
			continue
		}
		newQuery = append(newQuery, bm)
	}
	return newQuery, nil
}

func getDataBase(url string) string {
	index := strings.LastIndex(url, "/")
	return url[index+1:]
}

func Collection(tableName string) *mongo.Collection {
	return MongodbClient().Database(DatabaseName).Collection(tableName)
}

func MongodbConn() *mongo.Database {
	return MongodbClient().Database(DatabaseName)
}

func MongodbClient() *mongo.Client {
	return client
}

func CloseCursor(cursor *mongo.Cursor) {
	if cursor != nil {
		err := cursor.Close(context.TODO())
		if err != nil {
			fmt.Printf("mongodb cursor error %v \n", err)
		}
	}
}

type TotalCount struct {
	Id    string
	Total int64 `json:"total" bson:"total"`
}
