package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu         sync.Mutex
	cacheEntry map[string]cacheEntry
	interval   time.Duration
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:         sync.Mutex{},
		cacheEntry: make(map[string]cacheEntry),
		interval:   interval,
	}
	go cache.reapLoop()
	return cache
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	value, ok := c.cacheEntry[key]
	c.mu.Unlock()
	if ok {
		return value.val, true
	} else {
		return nil, false
	}
}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key := range c.cacheEntry {
			if time.Since(c.cacheEntry[key].createdAt) > c.interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
