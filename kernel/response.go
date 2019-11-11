package kernel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func (this Response) Success(data interface{}) {
	result := map[string]interface{}{}
	result["code"] = 0;
	result["data"] = data;
	this.Ctx.JSON(http.StatusOK, result)
}

func (this Response) Failed(code int, message string) {
	result := map[string]interface{}{}
	result["code"] = 0;
	result["message"] = message;
	this.Ctx.JSON(http.StatusOK, result)
}
