package commands

import "fmt"

func inspectCommand(cmdConfig *Config) error {

	pokemon, found := cmdConfig.Pokedex.Get(cmdConfig.PokemonToInspect)

	if found {
		pokemon.Inspect()
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
