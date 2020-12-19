package ws

import (
	"encoding/json"
	"os"
	"pervasive-chain/log"
	"sync"
	"time"
)

type ClientManager struct {
	Clients      map[*Client]bool
	Broadcast    chan interface{}
	BroadcastAll chan []byte
	Register     chan *Client
	Unregister   chan *Client
	WsDispatch   WsDispatch
	CacheMessage []interface{}
	sync.Mutex
}

var Manager = &ClientManager{
	Broadcast:    make(chan interface{}, 1040), // todo 足够大？
	BroadcastAll: make(chan []byte, 100),       // todo 足够大？
	Register:     make(chan *Client, 100),
	Unregister:   make(chan *Client, 100),
	Clients:      make(map[*Client]bool),
}

func (manager *ClientManager) RegisterRouter(ws WsDispatch) {
	manager.WsDispatch = ws
}

func (manager *ClientManager) Start(c chan os.Signal) {

	go manager.collectCacheMsg()

	for {
		select {
		case <-c:
			manager.ClosetAllClient()
			return
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			log.Info("ws client conn ... ", conn.ID, conn.ClientIp, " total conn: ", len(manager.Clients))
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
			}
			log.Info("ws  client exit .... ", conn.ID, conn.ClientIp, "total conn: ", len(manager.Clients))
		case message := <-manager.BroadcastAll:
			for conn := range manager.Clients {
				if conn.CanPush {
					conn.Send <- message
				}
			}
		default:

		}
	}
}

func BroadcastMessage(msg interface{}) {
	Manager.Broadcast <- msg
}



func (manager *ClientManager) collectCacheMsg() {
	canSend := true
	timer := time.NewTicker(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			canSend = false
			subscribeResp := NewSubscribeResp(manager.CacheMessage)
			if len(manager.CacheMessage) == 0 {
				log.Debug("send cache message is empty,continue ")
				continue
			}
			bytes, err := json.Marshal(subscribeResp)
			log.Debug("send subscribe info:  ", string(bytes))
			if err != nil {
				log.Error("send subscribe info: ", err.Error())
				return
			}
			manager.BroadcastAll <- bytes
			manager.CacheMessage = manager.CacheMessage[:0]
			canSend = true
		default:

		}
		if canSend {
			select {
			case msg := <-manager.Broadcast:
				manager.Lock()
				manager.CacheMessage = append(manager.CacheMessage, msg)
				manager.Unlock()
			default:

			}
		}
	}
}

func (manager *ClientManager) ClosetAllClient() {
	for client := range manager.Clients {
		manager.Unregister <- client
	}
}

func (manager *ClientManager) HeartBeat() {
	for conn := range manager.Clients {
		err := conn.Ping()
		if err != nil {
			manager.Unregister <- conn
		}
	}
}
