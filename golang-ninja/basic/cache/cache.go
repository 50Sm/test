package cache

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c Cache) Get(key string) int {
	value, exists := c.data[key]
	if exists {
		if intValue, ok := value.(int); ok {
			return intValue
		}
	}
	return 0
}

func (c Cache) Delete(key string) {
	delete(c.data, key)

}
