package kernel

import (
	"app/kernel/contract"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)

type RedisProvider struct {
}

func (this RedisProvider) Invoke() contract.DefinitionHandler {
	return func(container contract.ContainerInterface) interface{} {
		config := container.Get("config").(contract.ConfigInterface);
		addr := fmt.Sprintf("%s:%s", config.Get("REDIS_HOST", "localhost"), config.Get("REDIS_PORT", "6379"))
		password := config.Get("REDIS_AUTH", "")
		db, err := strconv.ParseInt(config.Get("REDIS_DB", "0"), 10, 64)
		if (err != nil) {
			db = 0;
		}

		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       int(db),
		})

		return client
	}
}
