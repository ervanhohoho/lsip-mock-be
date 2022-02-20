package main

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/controller"
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
	controller.InitHospitalController(r, &accessor)
	controller.InitUtilController(r, &accessor)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
