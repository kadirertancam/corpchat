package chat

import (
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" 
)

import "github.com/kadirertancam/corpchat/backend/internal/auth"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
 
func WsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"}); return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil { return }

		client := &Client{
			conn:   conn,
			userID: int64(claims.UserID),
			send:   make(chan []byte, 256),
			hub:    hub,
		}
		hub.register <- client

		go client.writePump()
		client.readPump()
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil { return }
		var msg Message
		if json.Unmarshal(msgBytes, &msg) != nil { continue }
		msg.FromUID = c.userID
		c.hub.broadcast <- msg
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, msg)
	}
}