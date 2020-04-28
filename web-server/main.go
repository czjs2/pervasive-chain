package main

import (
	"log"
	"pervasive-chain/config"
	"pervasive-chain/db"
	"pervasive-chain/httpsvr"
	lg "pervasive-chain/log"
)

func main() {
	prjConfig, err := config.ReadCfg("./web-config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = lg.MyLogicLogger(prjConfig.LogPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.InitMongo(prjConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = httpsvr.ListenAndServe(prjConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
