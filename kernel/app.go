package kernel

import (
	"app/kernel/contract"
)

var DI Container
var providers map[string]interface{}

func init() {
	DI = GetContainer()
	providers = map[string]interface{}{}

	providers["config"] = ConfigProvider{}
	providers["db"] = GormProvider{}
	providers["redis"] = RedisProvider{}
}

func BootStrap() {
	for key, value := range providers {
		if ret, ok := value.(contract.ProviderInterface); ok {
			DI.SetDefinition(key, ret.Invoke())
		}
	}
}
