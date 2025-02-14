package pokedex

import (
	"fmt"
	"pokedexcli/internal/pokemonapi"
)

type Pokedex struct {
	pokedex map[string]pokemonapi.Pokemon
}

func (dex *Pokedex) Add(pokemonName string, pokemon pokemonapi.Pokemon) {
	dex.pokedex[pokemonName] = pokemon
}

func (dex *Pokedex) Get(pokemonName string) (pokemonapi.Pokemon, bool) {
	pokemon, found := dex.pokedex[pokemonName]
	return pokemon, found
}

func (dex *Pokedex) DisplayPokedex() {
	fmt.Println("Your Pokedex:")
	for key, _ := range dex.pokedex {
		fmt.Printf(" - %v\n", key)
	}
}

func NewPokedex() Pokedex {
	dex := Pokedex{}
	dex.pokedex = make(map[string]pokemonapi.Pokemon)
	return dex
}
