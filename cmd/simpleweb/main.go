package main

import (
	"flag"
	"net"
	"net/http"
	"time"
)

var host string

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	flag.StringVar(&host, "host", ":8899", "web host addr")
	flag.Parse()

	server := &http.Server{
		Handler:      &Handler{},
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	listen, err := net.Listen("tcp4", host)
	if err != nil {
		panic(err)
	}
	server.SetKeepAlivesEnabled(false)
	err = server.Serve(listen)
}
