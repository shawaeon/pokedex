package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {	
	scanner := bufio.NewScanner(os.Stdin)	
	startURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	
	cfg := Config {
		Next: &startURL,
	}

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
			err := command.callback(&cfg)
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
	callback	func(cfg *Config) error
}

// Keeps track of location
type Config struct {
	Next 		*string
	Previous	*string
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
		"map": {
			name: 			"map",
			description:	"Display next 20 locations",
			callback: 		commandMapForward,
		},"mapb": {
			name: 			"mapb",
			description:	"Display previous 20 locations",
			callback: 		commandMapBack,
		},
	}
}