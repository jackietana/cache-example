package cache

import (
	"sync"
	"time"
)

type Value struct {
	value      any
	ttl        time.Duration
	expiration time.Time
}

type Cache struct {
	data  map[string]Value
	mutex *sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data:  make(map[string]Value),
		mutex: new(sync.RWMutex),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	c.data[key] = Value{
		value, ttl, time.Now().Add(ttl),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, exists := c.data[key]
	if !exists {
		return nil
	}

	if time.Now().Before(value.expiration) || time.Now().Equal(value.expiration) {
		return value.value
	}

	return nil
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	delete(c.data, key)
	c.mutex.Unlock()
}
