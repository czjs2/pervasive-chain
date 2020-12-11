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

	err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"event":"block","body":{},"msgId":"msgId%d"}`, time.Now().Unix())))
	if err != nil {
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

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			//chainInfoCmd := fmt.Sprintf(`{"uri":"chainInfo","body":{},"msgId":"msgId%d"}`, time.Now().Unix())
			//err = c.WriteMessage(websocket.TextMessage, []byte(chainInfoCmd))
			//if err != nil {
			//	fmt.Println(err.Error())
			//	return
			//}


			time.Sleep(3*time.Second)
			blockInfoCmd := fmt.Sprintf(`{"uri":"blockInfo","body":{"type":"S","chainKey":"SFF01","height":1,"hash":"sdfsdfs"},"msgId":"msgId%d"}`, time.Now().Unix())
			err = c.WriteMessage(websocket.TextMessage, []byte(blockInfoCmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}


			//time.Sleep(3*time.Second)
			//cmdInfo := fmt.Sprintf(`{"uri":"cmd","body":{"type":"S","cmd":{"key":"transfer","params":{"amount":100}}},"msgId":"msgId%d"}`, time.Now().Unix())
			//err = c.WriteMessage(websocket.TextMessage, []byte(cmdInfo))
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}
			//
			//
			//time.Sleep(3*time.Second)
			//ssInfo := fmt.Sprintf(`{"uri":"ssInfo","body":{"height":1,"fromShard":"SFF01","toShard":"S0000"}},"msgId":"msgId%d"}`, time.Now().Unix())
			//err = c.WriteMessage(websocket.TextMessage, []byte(ssInfo))
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}
			//
			//
			//time.Sleep(3*time.Second)
			//transInfo := fmt.Sprintf(`{"uri":"tranInfo","body":{"hash":"sdfsdf"}},"msgId":"msgId%d"}`, time.Now().Unix())
			//err = c.WriteMessage(websocket.TextMessage, []byte(transInfo))
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}

		}
	}

}
