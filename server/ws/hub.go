package ws

import (
	"sync"
)

type Hub struct {
	Clients    map[*Client]struct{}
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]struct{}), // the clients connected with the server
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

var once sync.Once
var singleton *Hub

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// 如果有 client 注册，则将该 client 加入 clients 中
			h.Clients[client] = struct{}{}
		case client := <-h.Unregister:
			// 如果有 client 注销，则从 clients 中删除 client ，并且关闭其 send channel
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			// 消息广播
			// 如果有 message, 则遍历 clients 将 message 发送给它们的 send channel
			// 如果 client.Send 存在，则发送，否则关闭该 client.Send 并且删除该 client
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
