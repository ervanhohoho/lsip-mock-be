package main

import (
	"github.com/ervanhohoho/lsip-mock-be/accessor"
	"github.com/ervanhohoho/lsip-mock-be/controller"
	docs "github.com/ervanhohoho/lsip-mock-be/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	accessor := accessor.Initialize("103.15.172.2", "primanet_lsip_mock", "primanet_lsip_user", "lsipmock2021")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controller.InitHospitalController(r, &accessor)
	controller.InitUtilController(r, &accessor)
	docs.SwaggerInfo_swagger.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://internship-kemenkes.vercel.app"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
