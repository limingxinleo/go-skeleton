package contract

type ContainerInterface interface {
	Get(key string, value interface{}) interface{}
	Set(key string, value interface{})
	Has(key string) bool
}
