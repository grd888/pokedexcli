package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// NewCache creates a new Cache
func NewCache() *Cache {	
  c := &Cache{
		cache: make(map[string]cacheEntry),
	}
  c.reapLoop(10 * time.Minute)
	return c
}

// Add adds a value to the cache with the given key
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

// Get retrieves a value from the cache for the given key
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	
	return entry.val, true
}

// ReapLoop removes entries older than the given duration
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reap(interval)
		}
	}()
}

// reap removes entries older than the given duration
func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	for key, entry := range c.cache {
		if time.Since(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
