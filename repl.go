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

		// If input is a command call function belonging to it
		commandName := words[0]
		if command, exists := getCommands()[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unkown command")
		}
	}
}

// Convert to lowercase and split based on whitespaces
func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

// Available commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 			"help",
			description:	"Display a help message",
			callback: 		commandHelp,
		},
		"exit": {
			name: 			"exit",
			description:	"Exit the Pokedex",
			callback: 		commandExit,
		},		
	}
}