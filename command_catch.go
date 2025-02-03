package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokeball"
	"strings"
)

func commandCatch (cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("Please input pokemon to catch (f.ex. \"catch pikachu\").")
		return nil
	}
	url := fmt.Sprintf("%v/pokemon/%s",pokeapi.BaseURL, args[0])
	pokemon := pokeball.Pokemon{}
	
	err := pokeapi.GetData(&url, cfg.apiClient, cfg.cache, &pokemon)
	if err != nil {
		if strings.Contains(err.Error(), "unmarshalling"){
			fmt.Println("Invalid pokemon name")
			return nil
		}
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	fmt.Printf("%s caught!\n", pokemon.Name)
	cfg.pokeball.Add(pokemon.Name, pokemon)
		
	return nil
}