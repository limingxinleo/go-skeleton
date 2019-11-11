package contract

type HandlerFunc func(ContainerInterface) interface{}

type ContainerInterface interface {
	Get(key string, value interface{}) interface{}
	Set(key string, handler interface{})
	Has(key string) bool
}
