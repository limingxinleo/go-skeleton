package kernel

import (
	"app/kernel/contract"
)

var DI Container
var providers map[string]interface{}

func init() {
	DI = Container{Container: map[string]interface{}{}}
	providers = map[string]interface{}{}

	providers["config"] = ConfigProvider{}
}

func BootStrap() {
	for key, value := range providers {
		if ret, ok := value.(contract.ProviderInterface); ok {
			DI.Set(key, ret.Invoke())
		}
	}
}
