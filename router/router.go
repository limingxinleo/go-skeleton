package router

import (
	"app/controller"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", controller.IndexController{}.Index())
	r.GET("/orm", controller.OrmController{}.Index())
	return r
}
