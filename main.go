package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		userInput := scanner.Text()

		if res {
			commands := strings.Fields(strings.ToLower(userInput))
			command := commands[0]
			fmt.Printf("Your command was: %s\n", command)
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