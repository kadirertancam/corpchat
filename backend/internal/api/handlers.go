package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"corpchat/internal/auth"
)

type registerReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req registerReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
		}
		hash, _ := auth.HashPassword(req.Password)
		var id int
		err := db.QueryRow(`INSERT INTO users(username, password_hash) VALUES($1, $2) RETURNING id`,
			req.Username, hash).Scan(&id)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "username taken"}); return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
		}
		var hash string
		var uid int
		if err := db.QueryRow(`SELECT id, password_hash FROM users WHERE username=$1`, req.Username).Scan(&uid, &hash); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"}); return
		}
		if !auth.CheckPassword(hash, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"}); return
		}
		token, _ := auth.GenerateToken(uid, req.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}