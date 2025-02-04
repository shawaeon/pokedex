package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries		map[string]cacheEntry
	interval	time.Duration
	mu 			*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache {
		entries:	map[string]cacheEntry{},
		interval:	interval,
		mu:			&sync.Mutex{},
	}
	go newCache.reapLoop()
	return &newCache 
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry {
		createdAt:	time.Now(),
		val:		val,
	}	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	return entry.val, exists	
}

// Deletes entries older than interval from cache
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		tickTime := <- ticker.C
		c.mu.Lock()
		for key, entry := range(c.entries) {
			if entry.createdAt.Before(tickTime.Add(-1 * c.interval)) {								
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}