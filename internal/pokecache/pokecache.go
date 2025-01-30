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

func NewCache(intervalInSeconds int) *Cache {
	newCache := Cache {
		interval: time.Duration(intervalInSeconds) * time.Second,
	}
	newCache.reapLoop()
	return &newCache 
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry {
		createdAt:	time.Now(),
		val:		val,
	}	
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false		
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
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