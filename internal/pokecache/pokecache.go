package pokecache

import (
	"time"
)

// This function will create a new cache, it will return a pointer to the cache
// and create a new map with the zero value for it's type
func NewCache(interval time.Duration) *Cache {
	return &Cache{
		data: make(map[string]cacheEntry),
	}
}

// This function will add a key to the cache
// It will use a mutex to lock the cache while we are adding a new key
// It will defer the unlock
func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// This function will return a key and its value from the cache if it exists
// It will use a mutex to lock the cache while we are getting a key
// It will defer the unlock
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	// Check if the element exists
	elem, ok := cache.data[key]
	return elem.val, ok

}

func (cache *Cache) reapLoop() {
	return
}
