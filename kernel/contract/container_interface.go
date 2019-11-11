package contract

type DefinitionHandler func(ContainerInterface) interface{}

type ContainerInterface interface {
	Get(key string) interface{}
	Make(key string) interface{}
	Set(key string, handler interface{}) interface{}
	SetDefinition(key string, handler DefinitionHandler)
	Has(key string) bool
}
