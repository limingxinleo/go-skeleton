package provider

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)

var Redis *redis.Client

func init() {
	addr := fmt.Sprintf("%s:%s", Config.Get("REDIS_HOST", "localhost"), Config.Get("REDIS_PORT", "6379"))
	password := Config.Get("REDIS_AUTH", "")
	db, err := strconv.ParseInt(Config.Get("REDIS_DB", "0"), 10, 64)
	if err != nil {
		db = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       int(db),
	})

	Redis = client
}
