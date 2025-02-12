package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cachEntry
	mu           sync.RWMutex
}

type cachEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration, enableReap bool) *Cache {
	newCache := Cache{}
	newCache.cacheEntries = make(map[string]cachEntry)
	if enableReap {
		go reapLoop(&newCache, interval)
	}
	return &newCache
}
