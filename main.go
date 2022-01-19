package main

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/gin-gonic/gin"
)

func main() {
	accessor.Test()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
