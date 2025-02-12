package pokecache

func Get(cache *Cache, key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	entry, found := cache.cacheEntries[key]

	if found {
		return entry.val, true
	}

	return nil, false
}
