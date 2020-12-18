package main

import (
	"log"
	"pervasive-chain/config"
	"pervasive-chain/mongodb"
)

func main() {
	manager := mongodb.NewTableManager()
	prjConfig, err := config.ReadWebCfg("./web-config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = mongodb.MongodbInit(prjConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = manager.ReadCfg("./tablecfg.json")
	if err != nil {
		log.Fatal(err)
	}
	err = manager.Run()
	if err != nil {
		log.Fatal(err)
	}

}
