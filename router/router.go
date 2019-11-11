package router

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(context *gin.Context) {
		controller.IndexController{}.Index(context)
	})
	return r
}
