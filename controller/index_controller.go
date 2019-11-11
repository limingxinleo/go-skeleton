package controller

import "github.com/gin-gonic/gin"

type IndexController struct {
	Ctx *gin.Context
}

func (this IndexController) Index() {
	this.Ctx.JSON(0, "Hello World.")
}
