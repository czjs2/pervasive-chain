package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"pervasive-chain/form"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:9999", "http service address")
var host = "http://127.0.0.1:9999"

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	interrupt1 := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt1, os.Interrupt)
	//go socketClient(interrupt)
	go apiClient(interrupt1)
	msg := make(chan string, 1)
	<-msg
}

func apiClient(interrupt chan os.Signal) {





	heartForm := form.HeartBeatFrom{
		Type:   "b",
		Number: "100",
		Id:     "101",
		Time:   "2020-11-11 11:11:11",
	}

	reportBlockForm := form.ReportBlockForm{
		Type:     "b",
		Number:   "1000",
		Id:       "1000",
		Height:   100,
		Father:   "100",
		Hash:     "100",
		Vrf:      "100",
		Time:     "2020-11-11 11:11:11",
		Interval: 100,
		Trans:    100,
		Size:     100,
		Detail:   nil,
	}

	reportFlowForm := form.ReportFlowForm{
		Type:   "b",
		Number: "100",
		Id:     "100",
		Time:   "100",
		In:     100,
		Out:    100,
	}

	ticker := time.NewTicker(5* time.Second)
	defer ticker.Stop()
	for {
		select {

		case <-ticker.C:

			resp1, err := HeartBeat(host, "/v1.0/headbeat", "", heartForm)
			fmt.Println("heartbeat:  response",resp1, err)
			resp2, err := ReportBlock(host, "/v1.0/block", "", reportBlockForm)
			fmt.Println("reportBlock:  response",resp2, err)
			resp3, err := ReportFlow(host, "/v1.0/flow", "", reportFlowForm)
			fmt.Println("reportFlow:   ",resp3, err)

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
