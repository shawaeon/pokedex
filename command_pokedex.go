package main

import (
	"fmt"
)

// Print caught pokemon
func commandPokedex (cfg *Config, args ...string) error {
	caughtPokemon := cfg.pokeball.GetAll()
	if len(caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon yet.")
		return nil
	}
	fmt.Println("Your pokedex:")
	for name := range(caughtPokemon) {
		fmt.Printf("  -%s\n", name)
	}
	return nil
}