package pokecache

import "time"

func Add(cache *Cache, key string, val []byte) {
	cache.mu.Lock()
	cache.cacheEntries[key] = cachEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.mu.Unlock()
}