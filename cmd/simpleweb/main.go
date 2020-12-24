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
		writer.WriteHeader(http.StatusOK)
	})
	fmt.Printf("web start %v \n", time.Now())
	err := http.ListenAndServe(host, nil)
	if err != nil {
		panic(err)
	}
}
