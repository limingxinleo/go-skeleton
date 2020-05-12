package controller

import (
	"app/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrmController struct {
}

func (this OrmController) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := bootstrap.Container.Get("db").(*gorm.DB)
		db.Where("id = ?", 1).First()
	}
}
