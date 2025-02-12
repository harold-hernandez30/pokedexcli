package pokecache

import (
	"fmt"
	"time"
)

func reapLoop(cache *Cache, interval time.Duration) {

	ticker := time.NewTicker(interval)

	for {
		currentTime := <-ticker.C
		for key, cachEntry := range cache.cacheEntries {
			diff := currentTime.Sub(cachEntry.createdAt)
			if diff.Abs().Milliseconds() > interval.Abs().Milliseconds() {
				fmt.Printf("\n%d: Reaping: %v", diff, key)
				delete(cache.cacheEntries, key)
			}
		}
	}
}
