package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var host string

func main() {
	flag.StringVar(&host, "host", ":8899", "web host addr")
	flag.Parse()

	http.HandleFunc("/api/v1.0/block", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(`{"code": 200,"message": "OK","data": {}}`))
		if err != nil {
			fmt.Printf("write block %v \n", err)
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Printf("update block incomming  %v \n", time.Now())
	})
	fmt.Printf("web start %v \n", time.Now())
	err := http.ListenAndServe(host, nil)
	if err != nil {
		panic(err)
	}
}
