package provider

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		Config.Get("DB_USERNAME", "root"),
		Config.Get("DB_PASSWORD", ""),
		Config.Get("DB_HOST", "localhost"),
		Config.Get("DB_PORT", "3306"),
		Config.Get("DB_DATABASE", "hyperf"),
		Config.Get("DB_CHARSET", "utf8mb4"),
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
