package controller

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/model"
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

// GetHospitals GetAccessTimes godoc
// @Summary 	Get hospital list
// @Schemes
// @Description Get hospital list
// @Tags 		Hospital
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} model.HospitalListResponse
// @Router 		/hospitals/ [get]
func (controller *hospitalController) GetHospitals(c *gin.Context) {
	hospitals := controller.accessor.GetHospitals()
	res := model.HospitalListResponse{Hospitals: hospitals}
	c.JSON(200, res)
}

// UpdateHospital GetAccessTimes godoc
// @Summary 	Update hospital
// @Schemes
// @Description Update hospital
// @Tags 		Hospital
// @Accept 		json
// @Produce 	json
// @Param 		data 	body 	model.Hospital 	true 	"Update Hospital"
// @Success 	200 {object} model.CommonSuccessResponse
// @Router 		/hospitals/update [post]
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

// ReserveHospital GetAccessTimes godoc
// @Summary 	Reserve hospital
// @Schemes
// @Description Reserve hospital
// @Tags 		Hospital
// @Accept 		json
// @Produce 	json
// @Param 		data 	body 	model.Hospital 	true 	"Reserve Hospital"
// @Success 	200 {object} model.CommonSuccessResponse
// @Router 		/hospitals/reserve [post]
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
