package main

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/login", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := c.SaveUploadedFile(file, filepath.Join("uploads", file.Filename)); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
	})

	server.Run(":8080")
}
