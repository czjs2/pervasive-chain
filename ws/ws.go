package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"pervasive-chain/log"
	"pervasive-chain/model"
)

type Client struct {
	ID       string
	Socket   *websocket.Conn
	Send     chan []byte
	ClientIp string
	Dispatch *Dispatch
}

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var Manager = &ClientManager{
	Broadcast:  make(chan []byte, 100), // todo 足够大？
	Register:   make(chan *Client, 100),
	Unregister: make(chan *Client, 100),
	Clients:    make(map[*Client]bool),
}

//todo 安全配置相关

func (manager *ClientManager) Start(c chan os.Signal) {
	for {
		select {
		case <-c:
			manager.ClosetAllClient()
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			log.Logger.Infoln("ws client conn ...", conn.ID, conn.ClientIp, len(manager.Clients))
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
			}
			log.Logger.Infoln("ws  client exit ....", conn.ID, conn.ClientIp, len(manager.Clients))
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:

				}
			}
		}
	}
}

// 广播消息
func BroadcastBlock(msg interface{}) {
	bytes, err := NewSubscribeResp("block", []interface{}{msg})
	if err != nil {
		fmt.Println("NewSubscribeResp is error ", err.Error())
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

func (c *Client) Ping() error {
	if _, _, err := c.Socket.NextReader(); err != nil {
		return err
	}
	return nil
}

// 发送失败直接断掉重连
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("rec: ", string(message))
		cmd := model.Cmd{}
		err = json.Unmarshal(message, &cmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resp, err := c.Dispatch.Execute(cmd)
		if err != nil {
			// todo ?
			c.Send <- resp
		} else {
			c.Send <- resp
		}
	}
}

func (c *Client) Write() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
