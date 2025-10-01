package cache

import "fmt"

var (
	expiredItemError = fmt.Errorf("cache.ttl.expired")
	keyNotFound      = fmt.Errorf("cache.key.not_found")
)
