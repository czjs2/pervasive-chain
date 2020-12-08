package ws

import (
	"encoding/json"
	"os"
	"pervasive-chain/log"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	WsDispatch WsDispatch
	cacheBlock []interface{}
}

var Manager = &ClientManager{
	Broadcast:  make(chan []byte, 100), // todo 足够大？
	Register:   make(chan *Client, 100),
	Unregister: make(chan *Client, 100),
	Clients:    make(map[*Client]bool),
}

func (manager *ClientManager) RegisterRouter(ws WsDispatch) {
	manager.WsDispatch = ws
}

//todo 安全配置相关

func (manager *ClientManager) Start(c chan os.Signal) {
	for {
		select {
		case <-c:
			manager.ClosetAllClient()
			return
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			log.Logger.Println("ws client conn ...", conn.ID, conn.ClientIp, len(manager.Clients))
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
			}
			log.Logger.Println("ws  client exit ....", conn.ID, conn.ClientIp, len(manager.Clients))
		case message := <-manager.Broadcast:
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
	// todo 缓存
	subscribeResp := NewSubscribeResp([]interface{}{msg})
	bytes, err := json.Marshal(subscribeResp)
	log.Debug("send subscribe info:  ", string(bytes))
	if err != nil {
		log.Error("send subscribe info: ", err.Error())
		return
	}
	Manager.Broadcast <- bytes
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
