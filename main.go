package main

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/model"
	"github.com/ervanhohoho/lsip-mock-be/payload"
	"github.com/gin-gonic/gin"
)

func setupHospitals(r *gin.Engine, accessor *accessor.Accessor) {
	r.GET("/hospitals", func(c *gin.Context) {
		hospitals := accessor.GetHospitals()
		c.JSON(200, gin.H{
			"hospitals": hospitals,
		})
	})
	r.POST("/hospitals/update", func(c *gin.Context) {
		var request payload.UpdateHospitalPayload
		if err := c.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
			c.Error(err)
		}
		success := accessor.UpdateHospital(request.Entities)
		if success {
			c.JSON(200, gin.H{"success": "Success update!"})
		} else {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
	})
	r.POST("/hospitals/reserve", func(c *gin.Context) {
		var request payload.ReserveHospitalPayload
		if err := c.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
			c.Error(err)
		}
		success, errMsg := accessor.ReserveHospital(request.Entity)
		if success {
			c.JSON(200, gin.H{"success": "Success update!"})
		} else {
			c.JSON(500, gin.H{"error": errMsg})
		}
	})
}
func setupUtils(r *gin.Engine, accessor *accessor.Accessor) {
	r.GET("/access_times", func(c *gin.Context) {
		obj, err := accessor.GetAccessTimes()
		if *err != nil {
			c.Error(*err)
		} else {
			c.JSON(200, gin.H{"data": obj})
		}
	})
	r.POST("/access_times/update", func(c *gin.Context) {
		var request model.AccessTime
		if err := c.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
			c.Error(err)
		}
		_, err := accessor.UpdateAccessTimes(request)
		if *err != nil {
			c.Error(*err)
		} else {
			c.JSON(200, gin.H{"success": "Success"})
		}
	})
}
func main() {
	accessor := accessor.Initialize("103.15.172.2", "primanet_lsip_mock", "primanet_lsip_user", "lsipmock2021")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	setupHospitals(r, &accessor)
	setupUtils(r, &accessor)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
