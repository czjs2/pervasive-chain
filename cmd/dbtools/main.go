package main

import "pervasive-chain/mongodb"

func main() {
	manager := mongodb.NewTableManager()
	err := manager.ReadCfg("./tablecfg.json")
	if err != nil {
		panic(err)
	}
	err = manager.Run()
	if err != nil {
		panic(err)
	}

}
