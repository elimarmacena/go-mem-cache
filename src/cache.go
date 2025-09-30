package cache

import (
	"fmt"
	"sync"
	"time"
)

type (
	// base structure for the cache usage.
	Cache struct {
		data   map[string]any
		signal sync.Mutex
	}
	// structure used to create item with life cycle.
	ExpireItem struct {
		data   any
		expiry time.Time
	}
)

func (c *Cache) SetWithTTL(key string, ttl time.Duration, value any) {
	item := ExpireItem{
		data:   value,
		expiry: time.Now().Add(ttl),
	}
	c.Set(key, item)
}

func (c *Cache) Set(key string, data any) {
	c.signal.Lock()
	c.data[key] = data
	c.signal.Unlock()
}

func (c Cache) Get(key string) any {
	if data, exists := c.data[key]; exists {
		return data
	}
	return nil
}

func (c *Cache) Delete(key string) error {
	c.signal.Lock()
	if _, exists := c.data[key]; !exists {
		return fmt.Errorf("key [%s] does not exists", key)
	}
	delete(c.data, key)
	c.signal.Unlock()
	return nil
}

func (c *Cache) Clear() {
	c.signal.Lock()
	c.data = nil
	c.signal.Unlock()
}

func (e ExpireItem) HasExpired() bool {
	return time.Now().After(e.expiry)
}
