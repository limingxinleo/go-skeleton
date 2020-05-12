package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func (r Response) Success(data interface{}) {
	result := map[string]interface{}{}
	result["code"] = 0
	result["data"] = data
	r.Ctx.JSON(http.StatusOK, result)
}

func (r Response) Failed(code int, message string) {
	result := map[string]interface{}{}
	result["code"] = code
	result["message"] = message
	r.Ctx.JSON(http.StatusOK, result)
}

func (r Response) Html(name string, data interface{}) {
	r.Ctx.HTML(http.StatusOK, name, data)
}
