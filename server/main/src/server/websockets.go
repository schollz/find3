package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

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

type Websockets struct {
	connections map[string]*websocket.Conn
	sync.Mutex
}

var (
	ws Websockets
)

func init() {
	ws.Lock()
	defer ws.Unlock()
	ws.connections = make(map[string]*websocket.Conn)
}

func wshandler(c *gin.Context) {
	family := strings.TrimSpace(c.DefaultQuery("family", ""))
	device := strings.TrimSpace(c.DefaultQuery("device", ""))
	if family == "" {
		c.String(http.StatusBadRequest, "need family")
		return
	} else if device == "" {
		c.String(http.StatusBadRequest, "need device")
		return
	}
	// TODO: validate one-time-pass (otp)
	// otp := c.DefaultQuery("otp", "")
	// if otp == "" {
	// 	return
	// }

	var w http.ResponseWriter = c.Writer
	var r *http.Request = c.Request

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	ws.Lock()
	ws.connections[family+"-"+device] = conn
	ws.Unlock()
	go sendOutLocation(family, device)
	// Listen to the websockets
	// for {
	// 	t, msg, err := conn.ReadMessage()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(t, msg)

	// 	newMsg, err := json.Marshal("hi")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	conn.WriteMessage(t, newMsg)
	// }
}

//  SendMessageOverWebsockets will send a message over the websockets
func SendMessageOverWebsockets(family string, device string, msg []byte) (err error) {
	ws.Lock()
	defer ws.Unlock()
	if _, ok := ws.connections[family+"-"+device]; ok {
		err = ws.connections[family+"-"+device].WriteMessage(1, msg)
	}
	return
}
