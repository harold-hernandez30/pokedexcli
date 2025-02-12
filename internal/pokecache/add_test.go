package pokecache_test

import (
	"bytes"
	"fmt"
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	fakeCache := pokecache.NewCache(time.Second, false)
	cases := []struct {
		key      string
		val      []byte
		expected []byte
	}{
		{
			key:      "pokemon-cli",
			val:      []byte("Pokemoncli"),
			expected: []byte("Pokemoncli"),
		},
	}

	failedCount := 0
	for i, c := range cases {
		pokecache.Add(fakeCache, c.key, c.val)

		fmt.Printf("\nTest %d of %d:\n", i+1, len(cases))
		var bufferExpected, bufferActual bytes.Buffer
		bufferExpected.Write(c.expected)

		cacheResult, ok := pokecache.Get(fakeCache, c.key)

		if ok {
			bufferActual.Write(cacheResult)
		}

		fmt.Printf("key: %v\n", c.key)
		fmt.Printf("expected: %v\n", bufferActual.String())
		fmt.Printf("actual: %v\n", bufferExpected.String())

		fmt.Println("")

		if bufferActual.String() != bufferExpected.String() {
			failedCount++
		}

		if failedCount > 0 {
			t.Fatalf("FAILED: %d mismatched", failedCount)
		}

	}

}
