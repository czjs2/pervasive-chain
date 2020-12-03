package main

import (
	"log"
	"pervasive-chain/app"
)

func main() {
	err := app.Run("./web-config.json")
	if err != nil {
		log.Fatal(err)
	}
}
