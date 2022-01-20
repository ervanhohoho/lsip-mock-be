package main

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/gin-gonic/gin"
)

func main() {
	accessor := accessor.Initialize("103.15.172.2", "primanet_lsip_mock", "primanet_lsip_user", "lsipmock2021")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hospitals", func(c *gin.Context) {
		hospitals := accessor.GetHospitals()
		c.JSON(200, gin.H{
			"hospitals": hospitals,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
