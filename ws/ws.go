package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

//todo 安全配置相关

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
			}
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
func BroadCast(msg []byte) {
	Manager.Broadcast <- msg
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
		cmd := BaseCmd{}
		err = json.Unmarshal(message, &cmd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resp, err := c.Dispatch.Execute(cmd)
		if err != nil {
			bytes, err := NewRespErr(cmd)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			c.Send <- bytes
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
