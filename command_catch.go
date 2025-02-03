package main

import (
	"fmt"
	"math/rand"
	"strings"

	"pokedex/internal/pokeapi"
	"pokedex/internal/pokeball"
)

func commandCatch (cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please input pokemon to catch (f.ex. \"catch pikachu\")")
	}
	url := fmt.Sprintf("%v/pokemon/%s",pokeapi.BaseURL, args[0])
	pokemon := pokeball.Pokemon{}
	
	err := pokeapi.GetData(&url, cfg.apiClient, cfg.cache, &pokemon)
	if err != nil {
		if strings.Contains(err.Error(), "unmarshalling"){
			return fmt.Errorf("invalid pokemon name")
		}
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if tryCatch(pokemon){
		fmt.Printf("%s was caught!\n", pokemon.Name)	
		cfg.pokeball.Add(pokemon.Name, pokemon)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}

// Chance of catching pokemon based on its base experience
func tryCatch (p pokeball.Pokemon) bool {
	chance := 1.0 - (float64(p.BaseExperience) / 700.0)
	return chance > rand.Float64()
}