package main

import (
	"log"
	"pervasive-chain/app"
)

func main() {
	err := app.Run("./web-config.json")
	//err := app.Run("./web-config-01.json")
	if err != nil {
		log.Fatal(err)
	}
}
