package pokecache

import (
	"time"
)

// NewCache - function will create a new cache, it will return a pointer to the cache
// and create a new map with the zero value for it's type
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}
	c.ticker = time.NewTicker(interval)
	go c.reapLoop()
	return c
}

// Add - function will add a key to the cache
// It will use a mutex to lock the cache while we are adding a new key
// It will defer unlocking
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

// Get - function will return a key and its value from the cache if it exists
// It will use a mutex to lock the cache while we are getting a key and defer the unlocking
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Check if the element exists
	elem, ok := c.data[key]
	return elem.val, ok

}

func (c *Cache) reapLoop() {
	for range c.ticker.C {
		c.mu.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
