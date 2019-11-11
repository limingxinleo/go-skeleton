package controller

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (this IndexController) Index(context *gin.Context) {
	context.JSON(0, "Hello World.")
}
