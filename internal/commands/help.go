package commands

import (
	"fmt"
)


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for command, obj := range registry {
		fmt.Printf("%v: %v\n", command, obj.description)
	}
	return nil
}