package config

type Config struct {
	items map[string]string
}

func New(items map[string]string) Config {
	return Config{items: items}
}

func (c Config) Get(key string, value string) string {
	ret, ok := c.items[key]
	if ok {
		return ret
	}

	return value
}

func (c Config) Has(key string) bool {
	_, ok := c.items[key]
	if ok {
		return true
	}

	return false
}
