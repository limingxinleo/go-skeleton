package router

import (
	"app/controller"
	"app/kernel"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(context *gin.Context) {
		controller.IndexController{kernel.Response{context}}.Index()
	})
	return r
}
