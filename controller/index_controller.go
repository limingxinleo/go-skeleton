package controller

import (
	"app/kernel/http"
	"app/kernel/provider"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (IndexController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		project := provider.Config.Get("APP_NAME", "Hyperf")

		message := fmt.Sprintf("Hello %s.", project)

		http.Response{Ctx: c}.Success(message)
	}
}
