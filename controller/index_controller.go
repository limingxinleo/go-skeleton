package controller

import (
	"app/kernel"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (this IndexController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		kernel.Response{context}.Success("Hello World.")
	}
}
