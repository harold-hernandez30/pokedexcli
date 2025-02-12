package commands

import (
	"fmt"
	"pokedexcli/internal/pokemonapi"
)

func mapbCommand(cmdConfig *Config) error {

	if len(cmdConfig.Previous) == 0 {
		cmdConfig.Previous = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	pokeAreas, err := pokemonapi.FetchLocationArea(cmdConfig.Previous, cmdConfig.PokeCache)

	if err == nil {
		if len(pokeAreas.Previous) > 0 {
			cmdConfig.Previous = pokeAreas.Previous
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
