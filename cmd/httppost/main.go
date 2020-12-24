package main

import (
	"fmt"
	"pervasive-chain/httpsvr"
	"time"
)

func main() {

	for i := 0; i < 10000; i++ {
		go func() {
			client := httpsvr.NewHttpClient()
			_, err := client.Post("http://172.16.7.182:8899/api/v1.0/block", "application/json", nil)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
	time.Sleep(1*time.Hour)
}
