package controller

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/payload"
	"github.com/gin-gonic/gin"
)

type hospitalController struct {
	accessor *accessor.Accessor
}

func InitHospitalController(r *gin.Engine, accessor *accessor.Accessor) {
	controller := hospitalController{accessor: accessor}
	h := r.Group("/hospitals")
	h.GET("/", controller.GetHospitals)
	h.POST("/update", controller.UpdateHospital)
	h.POST("/reserve", controller.ReserveHospital)
}

func (controller *hospitalController) GetHospitals(c *gin.Context) {
	hospitals := controller.accessor.GetHospitals()
	c.JSON(200, gin.H{
		"hospitals": hospitals,
	})
}
func (controller *hospitalController) UpdateHospital(c *gin.Context) {
	var request payload.UpdateHospitalPayload
	if err := c.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		c.Error(err)
	}
	success := controller.accessor.UpdateHospital(request.Entities)
	if success {
		c.JSON(200, gin.H{"success": "Success update!"})
	} else {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	}
}
func (controller *hospitalController) ReserveHospital(c *gin.Context) {
	var request payload.ReserveHospitalPayload
	if err := c.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		c.Error(err)
	}
	success, errMsg := controller.accessor.ReserveHospital(request.Entity)
	if success {
		c.JSON(200, gin.H{"success": "Success update!"})
	} else {
		c.JSON(500, gin.H{"error": errMsg})
	}
}
