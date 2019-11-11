package kernel

type Config struct {
	items map[string]string
}

func (this *Config) Get(key string, value string) string {
	ret, ok := this.items [ key ]
	if (ok) {
		return ret
	}

	return value
}

func (this *Config) Has(key string) bool {
	_, ok := this.items [ key ]
	if (ok) {
		return true
	}

	return false
}

func (this *Config) Init(items map[string]string) *Config {
	this.items = items
	return this
}
