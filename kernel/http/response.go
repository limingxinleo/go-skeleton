package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func (this Response) Success(data interface{}) {
	result := map[string]interface{}{}
	result["code"] = 0
	result["data"] = data
	this.Ctx.JSON(http.StatusOK, result)
}

func (this Response) Failed(code int, message string) {
	result := map[string]interface{}{}
	result["code"] = code
	result["message"] = message
	this.Ctx.JSON(http.StatusOK, result)
}

func (this Response) Html(name string, data interface{}) {
	this.Ctx.HTML(http.StatusOK, name, data)
}
