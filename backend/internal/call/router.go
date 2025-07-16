package call

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SignalMsg struct {
	Type string                 `json:"type"` // offer, answer, ice
	Data map[string]interface{} `json:"data"`
	To   int                    `json:"to"`
	From int                    `json:"from"`
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func WsCallHandler(hub *CallHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetInt("uid")
		conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
		client := &CallClient{uid: uid, conn: conn, hub: hub}
		hub.register <- client
		client.readLoop()
	}
}
