package config

import (
	"app/kernel/contract"
	"github.com/joho/godotenv"
	"log"
)

type ConfigProvider struct {
}

func (this ConfigProvider) Invoke() contract.DefinitionHandler {
	return func(container contract.ContainerInterface) interface{} {
		var items map[string]string
		items, err := godotenv.Read()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		return Config{items}
	}
}
