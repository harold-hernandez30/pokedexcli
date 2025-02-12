package commands

import (
	"pokedexcli/internal/pokecache"
)

type Config struct {
	Next      string
	Previous  string
	PokeCache *pokecache.Cache
}
