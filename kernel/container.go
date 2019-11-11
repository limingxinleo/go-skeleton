package kernel

type Container struct {
	Container map[string]interface{}
}

func (this *Container) Get(key string, value interface{}) interface{} {
	ret, ok := this.Container [ key ]
	if (ok) {
		return ret
	}

	return value
}

func (this *Container) Set(key string, value interface{}) {
	this.Container[key] = value
}

func (this *Container) Has(key string) bool {
	_, ok := this.Container [ key ]
	if (ok) {
		return true
	}

	return false
}
