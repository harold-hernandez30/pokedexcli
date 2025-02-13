package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/commands"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	commands.RegisterCommand("exit")
	commands.RegisterCommand("help")
	commands.RegisterCommand("map")
	commands.RegisterCommand("mapb")
	commands.RegisterCommand("explore")

	requestConfig := commands.Config{
		Next:                "",
		Previous:            "",
		PokeCache:           pokecache.NewCache(60*time.Second, true),
		CurrentLocationArea: "",
	}

	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		userInput := scanner.Text()

		if res {
			commandsSplice := strings.Fields(strings.ToLower(userInput))
			command := commandsSplice[0]

			if command == "explore" && len(commandsSplice) > 1 {
				requestConfig.CurrentLocationArea = commandsSplice[1]
			}

			cli, ok := commands.GetCliCommand(command)

			if ok {
				err := cli.Callback(&requestConfig)
				if err != nil {
					fmt.Printf("Error found: %v", err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}

func cleanInput(text string) []string {
	input := []string{}

	text = strings.TrimSpace(text)
	for _, word := range strings.Split(text, " ") {
		if len(strings.TrimSpace(word)) > 0 {
			input = append(input, word)
		}
	}

	return input
}
