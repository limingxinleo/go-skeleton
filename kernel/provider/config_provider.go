package provider

import (
	"app/kernel/config"
	"app/kernel/contract"
	"github.com/joho/godotenv"
	"log"
)

var Config contract.ConfigInterface

func init() {
	items, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = config.New(items)
}
