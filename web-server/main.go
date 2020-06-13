package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pervasive-chain/config"
	"pervasive-chain/db"
	"pervasive-chain/httpsvr"
	lg "pervasive-chain/log"
	"pervasive-chain/ws"
	"syscall"
	"time"
)

func main() {

	c := make(chan os.Signal)
	//c1 := make(chan os.Signal)
	signal.Notify(c,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, config.SIGUSR1, config.SIGUSR2)
	//signal.Notify(c1,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, config.SIGUSR1, config.SIGUSR2)

	prjConfig, err := config.ReadWebCfg("./web-config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = lg.MyLogicLogger(prjConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.InitMongo(prjConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//todo
	go ws.Manager.Start(c)

	err = httpsvr.ListenAndServe(prjConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func exitOs(){
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s Default Sig Exit ...\n", now)
	os.Exit(0)
}
