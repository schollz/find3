package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections map[string]*websocket.Conn

func init() {
	connections = make(map[string]*websocket.Conn)
}

func wshandler(c *gin.Context) {
	family := c.DefaultQuery("family", "")
	if family == "" {
		return
	}
	var w http.ResponseWriter = c.Writer
	var r *http.Request = c.Request

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	connections[family] = conn
	go tt()
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(t, msg)

		newMsg, err := json.Marshal("hi")
		if err != nil {
			panic(err)
		}
		conn.WriteMessage(t, newMsg)
	}
}

func tt() {
	time.Sleep(3 * time.Second)
	connections["zack"].WriteMessage(1, []byte("hi zack"))
}
