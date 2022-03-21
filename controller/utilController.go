package controller

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/model"
	"github.com/gin-gonic/gin"
)

type UtilController struct {
	accessor *accessor.Accessor
}

func InitUtilController(r *gin.Engine, accessor *accessor.Accessor) {
	utilController := UtilController{accessor: accessor}
	r.GET("/access_times", utilController.GetAccessTimes)
	r.POST("/access_times/update", utilController.UpdateAccessTimes)
}

// GetAccessTimes godoc
// @Summary 	Get site's access times
// @Schemes
// @Description Get enabled register and login time for user
// @Tags 		Access Time
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} model.AccessTime
// @Router 		/access_times/ [get]
func (u *UtilController) GetAccessTimes(c *gin.Context) {
	obj, err := u.accessor.GetAccessTimes()
	if *err != nil {
		c.Error(*err)
	} else {
		c.JSON(200, gin.H{"data": obj})
	}
}

// UpdateAccessTimes godoc
// @Summary 	Update site's access times
// @Schemes
// @Description Update enabled register and login time for user
// @Tags 		Access Time
// @Accept 		json
// @Produce 	json
// @Param 		data 	body 	model.AccessTime 	true 	"Update Access Time"
// @Success 	200 	object 	model.CommonSuccessResponse
// @Router 		/access_times/update [post]
func (u *UtilController) UpdateAccessTimes(c *gin.Context) {
	var request model.AccessTime
	if err := c.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		c.Error(err)
	}
	_, err := u.accessor.UpdateAccessTimes(request)
	if *err != nil {
		c.Error(*err)
	} else {
		c.JSON(200, gin.H{"success": "Success"})
	}
}
