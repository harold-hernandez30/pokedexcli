package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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