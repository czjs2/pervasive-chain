package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)



func TestRunCommand(t *testing.T){
	manager := NewTableManager()
	manager.test()
}


func TestTrancation(t *testing.T){
	dao := NewDaoWithTable("test")
	err := dao.UseSession(context.TODO(), func(sessionContext context.Context) error {
		_, err := dao.InsertOne(sessionContext, bson.M{"name": "xjrw"})
		if err!=nil{
			return err
		}
		return fmt.Errorf("test error ")
	})
	if err!=nil{
		fmt.Errorf(err.Error())
	}
}



func TestManagerTest(t *testing.T){
	manager:= NewTableManager()
	err = manager.ReadCfg("./tablecfg.json")
	if err!=nil{
		panic(err)
	}
	err = manager.CreateTable()
	if err!=nil{
		panic(err)
	}
	err = manager.CreateIndex()
	if err!=nil{
		panic(err)
	}
}
