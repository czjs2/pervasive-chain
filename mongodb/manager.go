package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"time"
)

type TableManger struct {
	mgo         *mongo.Database
	tables      []*TableInfo
}

func (tm *TableManger) test() {
	result := tm.mgo.RunCommand(context.Background(),bson.D{{"createUser", "test01"},
		{"pwd", "test01"}, {"roles", []bson.M{{"role": "readWrite","db":"pynxtest"}}}})
	if result.Err()!=nil{
		panic(result.Err())
	}
}

func (tm *TableManger) CreateTable() error {
	collectionNames, err := tm.mgo.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("get all collection error %v \n", err)
	}
	for i := 0; i < len(tm.tables); i++ {
		table := tm.tables[i]
		if !tm.TableIsExists(table.TableName, collectionNames) {
			err := tm.createTable(table.TableName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (tm *TableManger) CreateIndex() error {
	for i := 0; i < len(tm.tables); i++ {
		tableInfo := tm.tables[i]
		err := tm.createIndex(tableInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tm *TableManger) createIndex(tableInfo *TableInfo) error {
	collection := tm.mgo.Collection(tableInfo.TableName)
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		return err
	}
	// 获取表所有索引
	indexes, err := tm.getCollectionIndexes(cursor)
	if err != nil {
		return err
	}
	for i := 0; i < len(tableInfo.Indexes); i++ {
		param := tableInfo.Indexes[i]
		if !tm.indexExists(param, indexes) {
			// 创建索引
			// todo
			createIndexes := options.CreateIndexes()
			createIndexes.SetMaxTime(5 * time.Minute)
			_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: param})
			if err != nil {
				return fmt.Errorf("create index error %v %v \n", tableInfo.TableName, param)
			}
		}
	}
	return nil
}

func (tm *TableManger) indexExists(param param, indexes []*Index) bool {
	for i := 0; i < len(indexes); i++ {
		index := indexes[i]
		if isParamEqual(param, index.Key) { // 存在
			return true
		}
	}
	return false
}

func (tm *TableManger) getCollectionIndexes(cursor *mongo.Cursor) ([]*Index, error) {
	defer func() {
		if cursor != nil {
			_ = cursor.Close(context.TODO())
		}
	}()
	var res []*Index
	for cursor.Next(context.TODO()) {
		index := &Index{}
		err := cursor.Decode(index)
		if err != nil {
			return nil, err
		}
		res = append(res, index)
	}
	return res, nil
}

// 两个map中的元素是否相同
func isParamEqual(src, dest param) bool {
	if len(src) != len(dest) {
		return false
	}
	var total int
	for sk, sv := range src {
		ts := fmt.Sprintf("%v%v", sk, sv)
		for dk, dv := range dest {
			if ts == fmt.Sprintf("%v%v", dk, dv) {
				total = total + 1
			}
		}
	}
	return len(src) == total
}

func (tm *TableManger) createTable(table string) error {
	return tm.mgo.CreateCollection(context.TODO(), table)
}

func (tm *TableManger) TableIsExists(name string, collections []string) bool {
	for i := 0; i < len(collections); i++ {
		if collections[i] == name {
			return true
		}
	}
	return false
}

func (tm *TableManger) ReadCfg(path string) error {
	var res []*TableInfo
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read database table info %v \n", err)
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return fmt.Errorf("databse table info error %v \n", err)
	}
	tm.tables = res
	return nil
}

func NewTableManager() (*TableManger, ) {
	return &TableManger{
		mgo: MongodbConn(),
	}
}

type TableInfo struct {
	TableName string
	Indexes   []param
}

type param map[string]int

type Index struct {
	V    int    `bson:"v"`
	Key  param  `bson:"key"`
	Name string `bson:"name"`
	Ns   string `bson:"ns"`
}
