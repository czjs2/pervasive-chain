package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:8888", "http service address")

func main() {

	log.SetFlags(0)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/v1.0/wsConn"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	done := make(chan struct{})

	err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"uri":"event","body":{},"msgId":"msgId%d"}`, time.Now().Unix())))
	if err!=nil{
		log.Fatal(err)
	}

	defer c.Close()
	go func() {
		for {
			select {
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

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			chainInfoCmd := fmt.Sprintf(`{"uri":"chainInfo","body":{},"msgId":"msgId%d"}`, time.Now().Unix())
			err = c.WriteMessage(websocket.TextMessage, []byte(chainInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			blockInfoCmd := fmt.Sprintf(`{"uri":"blockInfo","body":{"type":"b","number":"0"},"msgId":"msgId%d"}`, time.Now().Unix())
			err = c.WriteMessage(websocket.TextMessage, []byte(blockInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			cmdInfo := fmt.Sprintf(`{"uri":"cmd","body":{"type":"b","cmd":{"key":"transfer","params":{"amount":100}}},"msgId":"msgId%d"}`, time.Now().Unix())
			err := c.WriteMessage(websocket.TextMessage, []byte(cmdInfo))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}

}
