package config

type Config struct {
	items map[string]string
}

func New(items map[string]string) Config {
	return Config{items: items}
}

func (this Config) Get(key string, value string) string {
	ret, ok := this.items[key]
	if ok {
		return ret
	}

	return value
}

func (this Config) Has(key string) bool {
	_, ok := this.items[key]
	if ok {
		return true
	}

	return false
}
