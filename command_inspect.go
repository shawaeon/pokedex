package main

import "fmt"

func commandInspect (cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please input pokemon to inspect (f.ex. \"inspect pikachu\")")
	}
	
	pokemon, exists := cfg.pokeball.Get(args[0])
	if !exists {
		return fmt.Errorf("pokemon not caught yet")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	
	fmt.Printf("Stats: \n")
	for _, stat := range(pokemon.Stats) {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")
	for _, typeInfo := range(pokemon.Types) {
		fmt.Printf("  -%v\n", typeInfo.Type.Name)
	}

	return nil
}