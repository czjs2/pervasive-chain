package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var host string

func main() {
	flag.StringVar(&host, "host", ":8899", "host addr ")
	flag.Parse()
	r := gin.Default()
	r.Handle("POST", "/xjrwTest", func(context *gin.Context) {
		fmt.Printf("gin  %v \n", time.Now())
	})
	r.Handle("GET", "/xjrwTest", func(context *gin.Context) {
		fmt.Printf("gin  %v \n", time.Now())
	})
	err := r.Run(host)
	if err != nil {
		panic(err)
	}
}
