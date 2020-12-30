package app

import (
	"fmt"
	"os"
	"os/signal"
	"pervasive-chain/config"
	"pervasive-chain/httpsvr"
	lg "pervasive-chain/log"
	"pervasive-chain/mongodb"
	"pervasive-chain/ws"
	"syscall"
	"time"
)

func Run(path string) error {
	c := make(chan os.Signal)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, config.SIGUSR1, config.SIGUSR2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, config.SIGUSR1, config.SIGUSR2)
	prjConfig, err := config.ReadWebCfg(path)
	if err != nil {
		return err
	}
	_, err = lg.MyLogicLogger(prjConfig.LogPath)
	if err != nil {
		return err
	}
	err = mongodb.MongodbInit(prjConfig)
	if err != nil {
		return err
	}
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Printf("app exit %v  %v \n", s, time.Now())
				exitOs()
			case config.SIGUSR1:
				fmt.Printf("app exit %v  %v \n", s, time.Now())
				exitOs()
			case config.SIGUSR2:
				fmt.Printf("app exit %v  %v \n", s, time.Now())
				exitOs()
			default:
				fmt.Printf("app exit %v  %v \n", s, time.Now())
			}
		}
	}()

	// todo
	ws.Manager.RegisterRouter(httpsvr.RegisterWsRouter())
	go ws.Manager.Start(c)

	err = httpsvr.ListenAndServe(prjConfig)

	return err
}

func exitOs() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s Default Sig Exit ...\n", now)
	os.Exit(0)
}
