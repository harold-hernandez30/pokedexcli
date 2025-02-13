package commands

import (
	"pokedexcli/internal/pokecache"
)

type Config struct {
	Next                string
	Previous            string
	CurrentLocationArea string
	PokeCache           *pokecache.Cache
}
