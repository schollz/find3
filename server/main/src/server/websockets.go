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
	connections map[string]map[string]*websocket.Conn
	sync.Mutex
}

var (
	ws Websockets
)

func init() {
	ws.Lock()
	defer ws.Unlock()
	ws.connections = make(map[string]map[string]*websocket.Conn)
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
	if _, ok := ws.connections[family+"-"+device]; !ok {
		ws.connections[family+"-"+device] = make(map[string]*websocket.Conn)
	}
	ws.connections[family+"-"+device][conn.RemoteAddr().String()] = conn
	ws.Unlock()
	go sendOutLocation(family, device)
	go websocketListener(family, device, conn)
	// Listen to the websockets

}

func websocketListener(family string, device string, conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			ws.Lock()
			if _, ok := ws.connections[family+"-"+device]; ok {
				if _, ok2 := ws.connections[family+"-"+device][conn.RemoteAddr().String()]; ok2 {
					delete(ws.connections[family+"-"+device], conn.RemoteAddr().String())
				}
				logger.Log.Debugf("removed %s/%s", family+"-"+device, conn.RemoteAddr().String())
			}
			ws.Unlock()
			return
		}
	}
}

//  SendMessageOverWebsockets will send a message over the websockets
func SendMessageOverWebsockets(family string, device string, msg []byte) (err error) {
	ws.Lock()
	defer ws.Unlock()
	if _, ok := ws.connections[family+"-"+device]; ok {
		for _, conn := range ws.connections[family+"-"+device] {
			err = conn.WriteMessage(1, msg)
			if err != nil {
				logger.Log.Warnf("problem sending websocket: %s/%s", family+"-"+device)
			}
		}
	}
	return
}
