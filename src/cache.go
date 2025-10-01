package cache

import (
	"sync"
	"time"
)

type (
	// base structure for the cache usage.
	CacheTTL struct {
		data   map[string]item
		signal sync.Mutex
	}
	// structure used to create item with life cycle.
	item struct {
		value  any
		expiry time.Time
	}
)

// ---- CacheTTL

// Set - Creates a new entry into the cache given the key, ttl and value.
//
// The operaion it is thread safe, once that uses operation lock.
func (c *CacheTTL) Set(key string, ttl time.Duration, value any) {
	item := item{
		value:  value,
		expiry: time.Now().Add(ttl),
	}

	c.signal.Lock()
	c.data[key] = item
	c.signal.Unlock()
}

// Get- searchs for key and verify the TTL.
// When TTL is expired, an error is raised.
func (c CacheTTL) Get(key string) (any, error) {
	if item, exists := c.data[key]; exists && item.HasExpired() {
		return item.value, nil
	} else if exists {
		return nil, expiredItemError
	}
	return nil, nil
}

// Delete - searchs a key into the cache and then remove the value.
// If no key is found, an error is raised.
//
// The operaion it is thread safe, once that uses operation lock.
func (c *CacheTTL) Delete(key string) error {
	c.signal.Lock()
	if _, exists := c.data[key]; !exists {
		return keyNotFound
	}
	delete(c.data, key)
	c.signal.Unlock()
	return nil
}

// Clear - removes all keys available in the cache.
//
// The operaion it is thread safe, once that uses operation lock.
func (c *CacheTTL) Clear() {
	c.signal.Lock()
	c.data = nil
	c.signal.Unlock()
}

// ---- item

// HasExpired - checks if the current TTL is expired.
func (e item) HasExpired() bool {
	return e.expiry.After(time.Now())
}
