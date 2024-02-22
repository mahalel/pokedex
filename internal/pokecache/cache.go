package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents an entry in the cache.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache is a struct that holds the cache data and provides methods for safe access.
type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}
