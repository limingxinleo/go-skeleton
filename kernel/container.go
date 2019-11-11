package kernel

import (
	"app/kernel/contract"
)

type Container struct {
	Container  map[string]interface{}
	Definition map[string]contract.DefinitionHandler
}

func GetContainer() Container {
	return Container{
		Container:  map[string]interface{}{},
		Definition: map[string]contract.DefinitionHandler{},
	}
}

func (this Container) Get(key string) interface{} {
	ret, ok := this.Container [ key ]
	if (ok) {
		return ret
	}

	return DI.Set(key, this.Make(key))
}

func (this Container) Make(key string) interface{} {
	ret, ok := this.Definition [ key ]
	if (!ok) {
		return nil
	}

	return ret(this)
}

func (this Container) Set(key string, value interface{}) interface{} {
	this.Container[key] = value
	return value
}

func (this Container) SetDefinition(key string, handler contract.DefinitionHandler) {
	this.Definition[key] = handler
}

func (this Container) Has(key string) bool {
	_, ok := this.Container [ key ]
	if (ok) {
		return true
	}

	return false
}
