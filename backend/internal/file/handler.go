package file

import (
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func UploadHandler(mc *minio.Client, bucket string) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		ext := filepath.Ext(header.Filename)
		objectName := "uploads/" + time.Now().Format("20060102/150405") + ext

		_, err = mc.PutObject(c, bucket, objectName, file, header.Size,
			minio.PutObjectOptions{ContentType: header.Header.Get("Content-Type")})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		url := "https://cdn.yourdomain.com/" + objectName
		c.JSON(200, gin.H{"url": url})
	}
}
