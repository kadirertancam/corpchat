package main

import (
    "log"
    "os"
     
    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"


)
import "github.com/kadirertancam/corpchat/backend/internal/db"
import "github.com/kadirertancam/corpchat/backend/internal/api"

import "github.com/kadirertancam/corpchat/backend/internal/chat"

func main() {
    dbx, err := sqlx.Connect("postgres", mustGetEnv("DB_DSN"))
    if err != nil {
        log.Fatal(err)
    }

   if err := db.Migrate(dbx.DB); err != nil {
    log.Fatal(err)
}
    hub := chat.NewHub()
    go hub.Run()

    r := gin.Default()
    r.GET("/ws", chat.WsHandler(hub))
    r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
    r.POST("/register", api.Register(dbx))
    r.POST("/login", api.Login(dbx))

    r.Run(":8080")
}

func registerHandler(db *sqlx.DB) gin.HandlerFunc {
    type req struct{ Username, Password string }
    return func(c *gin.Context) {
        var q req
        if err := c.BindJSON(&q); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        _, _ = db.Exec(`INSERT INTO users(username, password_hash) VALUES($1, crypt($2, gen_salt('bf')))`, q.Username, q.Password)
        c.JSON(201, gin.H{"ok": true})
    }
}

func mustGetEnv(key string) string {
    val := os.Getenv(key)
    if val == "" {
        log.Fatalf("Environment variable %s is required.", key)
    }
    return val
}
