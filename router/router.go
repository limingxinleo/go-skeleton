package router

import (
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.Default()

	return r
}
