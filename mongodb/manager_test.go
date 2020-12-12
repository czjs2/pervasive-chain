package mongodb

import "testing"



func TestRunCommand(t *testing.T){
	manager := NewTableManager()
	manager.test()
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
