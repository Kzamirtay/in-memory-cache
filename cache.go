package cache

type Cache struct {
	storage map[string]interface{}
}

func New() Cache {
	return Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) *Cache {
	c.storage[key] = value

	return c
}

func (c *Cache) Get(key string) interface{} {
	return c.storage[key]
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}
