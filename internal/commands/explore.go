package commands

import (
	"fmt"
	"pokedexcli/internal/pokemonapi"
)

func exploreCommand(cmdCommand *Config) error {

	if len(cmdCommand.CurrentLocationArea) == 0 {
		return fmt.Errorf("no location specified")
	}

	fmt.Printf("Exploring %v...\n", cmdCommand.CurrentLocationArea)

	res, err := pokemonapi.FetchPokemons(cmdCommand.CurrentLocationArea, cmdCommand.PokeCache)

	if err != nil {
		return fmt.Errorf("error fetching pokemons: %v", err)
	}

	for _, pokemonEncounter := range res.PokemonEncounters {
		fmt.Printf("%v\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
