package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kadirertancam/corpchat/backend/internal/api"
	"github.com/kadirertancam/corpchat/backend/internal/chat"
	"github.com/kadirertancam/corpchat/backend/internal/db"
	"github.com/kadirertancam/corpchat/backend/internal/file"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	dbx, err := sqlx.Connect("postgres", mustGetEnv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(dbx.DB); err != nil {
		log.Fatal(err)
	}
	hub := chat.NewHub(dbx)
	go hub.Run()
	minioClient, _ := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})

	r := gin.Default()
	r.Static("/cdn", "./uploads")
	r.POST("/upload", file.UploadHandler(minioClient, "corpchat"))
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
