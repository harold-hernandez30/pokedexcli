package commands

import (
	"fmt"
	"pokedexcli/internal/pokemonapi"
)

func mapCommand(cmdConfig *Config) error {

	if len(cmdConfig.Next) == 0 {
		cmdConfig.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	pokeAreas, err := pokemonapi.FetchLocationArea(cmdConfig.Next, cmdConfig.PokeCache)

	if err == nil {
		if len(pokeAreas.Next) > 0 {
			cmdConfig.Next = pokeAreas.Next
		}

		if len(pokeAreas.Previous) > 0 {
			cmdConfig.Previous = pokeAreas.Previous
		}

		for _, pokeArea := range pokeAreas.LocationAreas {
			fmt.Printf("%s\n", pokeArea.Name)
		}

		return nil
	} else {
		return fmt.Errorf("something went wrong: %v", err)
	}
}
