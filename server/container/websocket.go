package container

import (
	"log"
	"net/http"

	"github.com/fangaoxs/synk/server/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WsController(hub *ws.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		wshandler(hub, c.Writer, c.Request)
	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	// 将 http 升级为 ws 协议完成ws连接
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &ws.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
