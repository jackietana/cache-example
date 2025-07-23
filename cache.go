package cache

import "time"

type Value struct {
	value      any
	ttl        time.Duration
	expiration time.Time
}

type Cache struct {
	data map[string]Value
}

func New() *Cache {
	return &Cache{
		data: make(map[string]Value),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.data[key] = Value{
		value, ttl, time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) interface{} {
	curr := time.Now()
	expr := c.data[key].expiration

	if curr.Before(expr) || curr.Equal(expr) {
		return c.data[key].value
	} else {
		c.Delete(key)
		return nil
	}
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
