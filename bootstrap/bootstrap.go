package bootstrap

import (
	"app/kernel"
	"app/kernel/config"
	"app/kernel/contract"
	"app/kernel/orm"
	"app/kernel/redis"
)

var Container kernel.Container
var providers map[string]interface{}

func init() {
	Container = kernel.GetContainer()
	providers = map[string]interface{}{}

	providers["config"] = config.ConfigProvider{}
	providers["db"] = orm.GormProvider{}
	providers["redis"] = redis.RedisProvider{}
}

func BootStrap() {
	Container.Set("container", Container)
	for key, value := range providers {
		if ret, ok := value.(contract.ProviderInterface); ok {
			Container.SetDefinition(key, ret.Invoke())
		}
	}
}

func PreLoad() {
	for key, _ := range providers {
		Container.Get(key)
	}
}
