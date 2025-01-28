package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		fmt.Printf("Your command was: %v\n", command)
	}
}

// Convert to lowercase and split based on whitespaces
func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}