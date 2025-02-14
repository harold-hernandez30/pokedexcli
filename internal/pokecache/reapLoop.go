package pokecache

import (
	"time"
)

func reapLoop(cache *Cache, interval time.Duration) {

	ticker := time.NewTicker(interval)

	for {
		currentTime := <-ticker.C
		for key, cachEntry := range cache.cacheEntries {
			diff := currentTime.Sub(cachEntry.createdAt)
			if diff.Abs().Milliseconds() > interval.Abs().Milliseconds() {
				delete(cache.cacheEntries, key)
			}
		}
	}
}
