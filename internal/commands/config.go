package commands

import (
	"pokedexcli/internal/pokecache"
	"pokedexcli/internal/pokedex"
)

type Config struct {
	Next                string
	Previous            string
	CurrentLocationArea string
	PokemonToCatch      string
	PokemonToInspect    string
	Pokedex             pokedex.Pokedex
	PokeCache           *pokecache.Cache
}
