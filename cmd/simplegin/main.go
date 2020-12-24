package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var host string

func main() {
	flag.StringVar(&host, "host", ":8899", "host addr ")
	flag.Parse()
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Handle("POST", "/api/v1.0/block", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})
	r.Handle("GET", "/api/v1.0/block", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})
	err := r.Run(host)
	if err != nil {
		panic(err)
	}
}
