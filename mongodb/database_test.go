package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type Decimal struct {
	primitive.Decimal128
}

type Value struct {
	Value Decimal `json:"value" bson:"value"`
}

func Test001(t *testing.T) {
	//opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	//
	//idxFiles := []mongo.IndexModel{
	//	{
	//		Keys: bsonx.Doc{{"name": "text"}},
	//	},
	//}
	//
	//MongodbConn().Collection("mycollection").Indexes().CreateMany(context.TODO(), idxFiles, opts)
	//MongodbConn().Collection("mycollection").Indexes().List()
}



func TestCollection(t *testing.T) {
	dao := NewDaoWithTable("test")
	dd, err := primitive.ParseDecimal128("11111")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	_, err = dao.InsertOne(context.TODO(), bson.M{"value": dd})
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	value:=&Value{}
	_, err = dao.FindOne(context.TODO(), bson.M{"value": bson.M{"$exists": true}}, &value)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(value.Value.String())
}
