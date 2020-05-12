package controller

import (
	"app/kernel/http"
	"app/service/dao"
	"github.com/gin-gonic/gin"
)

type OrmController struct {
}

func (this OrmController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := dao.UserDao{}.First(1)

		http.Response{context}.Success(gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"gender":     user.Gender,
			"created_at": user.CreatedAt,
		})
	}
}
