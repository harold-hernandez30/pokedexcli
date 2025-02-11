package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/commands"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
		
	commands.RegisterCommand("exit")
	commands.RegisterCommand("help")
	commands.RegisterCommand("map")
	commands.RegisterCommand("mapb")

	requestConfig := commands.Config{
		Next: "",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		userInput := scanner.Text()

		if res {
			commandsSplice := strings.Fields(strings.ToLower(userInput))
			command := commandsSplice[0]

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

	return  input
}