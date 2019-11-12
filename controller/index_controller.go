package controller

import (
	"app/kernel/http"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (this IndexController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		http.Response{context}.Success("Hello World.")
	}
}
