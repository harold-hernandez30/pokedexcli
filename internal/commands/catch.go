package commands

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokemonapi"
)

func catchCommand(cmdConfig *Config) error {

	if len(cmdConfig.PokemonToCatch) == 0 {
		return fmt.Errorf("no pokemon specified")
	}

	pokemon, err := pokemonapi.FetchPokemon(cmdConfig.PokemonToCatch, cmdConfig.PokeCache)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if err != nil {
		return fmt.Errorf("error when fetching pokemon: %v", err)
	}

	chance := rand.Intn(400)

	if chance < pokemon.BaseExperience {

		fmt.Printf("%v was caught!\n", pokemon.Name)
		cmdConfig.Pokedex.Add(pokemon.Name, pokemon)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	cmdConfig.Pokedex.Add(pokemon.Name, pokemon)

	return nil
}
