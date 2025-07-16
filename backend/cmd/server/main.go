package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sqlx.Connect("postgres", mustGetEnv("DB_DSN"))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.POST("/register", registerHandler(db))
	r.Run(":8080")
}

func registerHandler(db *sqlx.DB) gin.HandlerFunc {
	type req struct{ Username, Password string }
	return func(c *gin.Context) {
		var q req
		if err := c.BindJSON(&q); err != nil {
			c.JSON(400, gin.H{"error": err.Error()}); return
		}
		_, _ = db.Exec(`INSERT INTO users(username, password_hash) VALUES($1, crypt($2, gen_salt('bf')))`, q.Username, q.Password)
		c.JSON(201, gin.H{"ok": true})
	}
}