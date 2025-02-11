package commands

import (
	"fmt"
	"os"
)


func commandExit(cliConfig *Config) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

