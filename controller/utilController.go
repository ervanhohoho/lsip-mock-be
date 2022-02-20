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
	r.GET("/access_times/update", utilController.UpdateAccessTimes)
}

func (u *UtilController) GetAccessTimes(c *gin.Context) {
	obj, err := u.accessor.GetAccessTimes()
	if *err != nil {
		c.Error(*err)
	} else {
		c.JSON(200, gin.H{"data": obj})
	}
}
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
