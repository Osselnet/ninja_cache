package cache

type cache struct {
	temple map[string]interface{}
}

func New() *cache {
	return &cache{
		temple: make(map[string]interface{}),
	}
}

func (c cache) Set(key string, value interface{}) {
	c.temple[key] = value
}

func (c cache) Get(key string) interface{} {
	return c.temple[key]
}

func (c cache) Delete(key string) {
	delete(c.temple, key)
}
