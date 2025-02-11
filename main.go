package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandsMap = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}

var registry = map[string]cliCommand{}

type cliCommand struct {
	name string
	description string
	callback func() error
}


func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		userInput := scanner.Text()

		registry["exit"] = commandsMap["exit"]
		registry["help"] = commandsMap["help"]

		if res {
			commands := strings.Fields(strings.ToLower(userInput))
			command := commands[0]

			cli, ok := commandsMap[command]

			if ok {
				err := cli.callback()
				if err != nil {
					fmt.Printf("Error found: %v", err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		} 
	}
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for command, obj := range registry {
		fmt.Printf("%v: %v\n", command, obj.description)
	}
	return nil
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