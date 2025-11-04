package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
		mu:      sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for now := range ticker.C {
		keysToDelete := []string{}
		c.mu.RLock()
		for key, value := range c.entries {
			// if the entry is older than the interval, delete it
			if value.createdAt.Add(interval).Before(now) {
				keysToDelete = append(keysToDelete, key)
			}
		}
		c.mu.RUnlock()

		if len(keysToDelete) == 0 {
			continue
		}

		c.mu.Lock()
		for _, key := range keysToDelete {
			delete(c.entries, key)
		}
		c.mu.Unlock()
	}
}
