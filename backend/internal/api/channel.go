package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kadirertancam/corpchat/backend/internal/chat"
)

var hub *chat.Hub // global pointer

func Init(h *chat.Hub) {
	hub = h
}

func PostChannelMessage(db *sqlx.DB) gin.HandlerFunc {
	type req struct {
		ChannelID int    `json:"channel_id" binding:"required"`
		Body      string `json:"body"`
	}

	return func(c *gin.Context) {
		var r req
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		uid := c.GetInt("uid")
		var msgID int64
		err := db.QueryRow(`
            INSERT INTO messages(from_uid, channel_id, body)
            VALUES($1,$2,$3) RETURNING id`,
			uid, r.ChannelID, r.Body).Scan(&msgID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Broadcast
		hub.BroadcastToChannel(r.ChannelID, chat.Message{
			ID:        msgID,
			FromUID:   int64(uid),
			ChannelID: &r.ChannelID,
			Body:      r.Body,
		})

		c.JSON(201, gin.H{"id": msgID})
	}
}
