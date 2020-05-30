package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:9999", "http service address")

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	interrupt1 := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt1, os.Interrupt)
	go socketClient(interrupt)
	go apiClient(interrupt1)
	msg:=make(chan string,1)
	<-msg
}

func apiClient(interrupt chan os.Signal) {

	heartBeat := `{
			   "type": "b"
			   "number": "123", 
			   "Id": "abcdid", 
			   "time":"2020-11-11: 11:11:11" 
    }`
	reportBlock := `{
			   "type": "b", 
			   "number": "123", 
			   "Id": "adfsf", 
				height: 111,
				father:"safdsfsdf",
				hash:"sfdsdfsf",
				vrf:"vrf",
				time:"2020-11-11 11:11:11",
				interval:45,
				trans:100,
				size: 400,
				detail:""
	}`

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {

		case <-ticker.C:
			_, _ = HeartBeat(*addr, "api/v1/headbeat", "", heartBeat)
			_, _ = ReportBlock(*addr, "api/v1/block", "", reportBlock)
		case <-interrupt:
			fmt.Println("httpClient interrupt ")
			return
		}
	}

}

func socketClient(interrupt chan os.Signal) {
	log.SetFlags(0)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/v1/conn"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {

			select {
			case <-interrupt:
				c.Close()
				fmt.Println(" websocket interrupt")
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
				return
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}
				log.Printf("recv: %s", message)
			}
		}
	}()
	ticker := time.NewTicker(2 * time.Second)

	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			res := `{"uri":"chainInfo","body":{},"msgId":"msgid111111111"}`
			err := c.WriteMessage(websocket.TextMessage, []byte(res))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
