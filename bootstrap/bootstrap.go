package bootstrap

import (
	"app/kernel"
	"app/kernel/contract"
)

var Container kernel.Container
var providers map[string]interface{}

func init() {
	Container = kernel.GetContainer()
	providers = map[string]interface{}{}

	providers["config"] = kernel.ConfigProvider{}
	providers["db"] = kernel.GormProvider{}
	providers["redis"] = kernel.RedisProvider{}
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
