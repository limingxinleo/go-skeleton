package contract

type ConfigInterface interface {
	Get(key string, value string) string
	Has(key string) bool
}
