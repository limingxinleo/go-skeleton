package dao

import (
	"app/bootstrap"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	DB = bootstrap.Container.Get("db").(*gorm.DB)
}
