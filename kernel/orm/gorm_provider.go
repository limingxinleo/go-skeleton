package orm

import (
	"app/kernel/contract"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormProvider struct {
}

func (this GormProvider) Invoke() contract.DefinitionHandler {
	return func(container contract.ContainerInterface) interface{} {
		config := container.Get("config").(contract.ConfigInterface);

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.Get("DB_USERNAME", "root"),
			config.Get("DB_PASSWORD", ""),
			config.Get("DB_HOST", "localhost"),
			config.Get("DB_PORT", "3306"),
			config.Get("DB_DATABASE", "hyperf"),
			config.Get("DB_CHARSET", "utf8mb4"),
		)

		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			panic("failed to connect database")
		}

		return db
	}
}
