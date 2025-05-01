package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")

	r.GET("/port", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port": port,
		})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": os.Getenv("test"),
		})
	})

	r.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
