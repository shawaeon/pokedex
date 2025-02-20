package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shawaeon/pokedex/internal/pokeapi"
	"github.com/shawaeon/pokedex/internal/pokeball"
	"github.com/shawaeon/pokedex/internal/pokecache"
)

func startRepl() {	
	scanner := bufio.NewScanner(os.Stdin)	
	startURL := pokeapi.BaseURL + "/location-area/?offset=0&limit=20"
	cacheInterval := 10 * time.Second
	
	cfg := Config {
		apiClient: 		&http.Client{Timeout: 5 * time.Second},
		cache:			pokecache.NewCache(cacheInterval),
		pokeball: 		pokeball.NewPokeball(),
		Next: 			&startURL,
	}

	commandHelp(&cfg)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		// location for commandExplore
		

		// If input is a command call function belonging to it
		commandName := words[0]
		if command, exists := getCommands()[commandName]; exists {
			var err error
			if len(words) > 1 {
				err = command.callback(&cfg, words[1])
			} else{
				err = command.callback(&cfg)
			}
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
	callback	func(cfg *Config, args ...string) error
}

// Location on the map, api client and cache
type Config struct {	
	apiClient		*http.Client
	cache			*pokecache.Cache
	pokeball		*pokeball.Pokeball
	Next 			*string
	Previous		*string
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
		},
		"mapb": {
			name: 			"mapb",
			description:	"Display previous 20 locations",
			callback: 		commandMapBack,
		},
		"explore": {
			name:			"explore <location>",
			description: 	"Explore a location",
			callback:		commandExplore,	
		},
		"catch": {
			name:			"catch <pokemon>",
			description: 	"Catch a pokemon",
			callback:		commandCatch,	
		},
		"inspect": {
			name:			"inspect <pokemon>",
			description:	"Get information about caught pokemon",
			callback: 		commandInspect,	
		},
		"pokedex": {
			name:			"pokedex",
			description: 	"Display all caught pokemon",
			callback: 		commandPokedex,
		},
	}
}