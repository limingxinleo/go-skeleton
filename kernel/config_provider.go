package kernel

import (
	"github.com/joho/godotenv"
	"log"
)

type ConfigProvider struct {
}

func (this ConfigProvider) Invoke() interface{} {
	var items map[string]string
	items, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config:= Config{}
	return config.Init(items)
}
