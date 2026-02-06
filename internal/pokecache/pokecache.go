package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	Entry    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entry: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.Entry[key]
	if !exists {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entry[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C{
		c.mu.Lock()
		for key, entry := range c.Entry {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.Entry, key)
			}
		}
		c.mu.Unlock()
	}
}
