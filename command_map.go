package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func commandMapForward(cfg *Config, args ...string) error {
	url := cfg.Next
	if url == nil {
		fmt.Println("No more locations")
		return nil
	} 
	return commandMap(cfg, url)
}

func commandMapBack(cfg *Config, args ...string) error {
	url := cfg.Previous
	if url == nil {
		fmt.Println("No previous locations")
		return nil
	} 
	return commandMap(cfg, url)
}

// Prints locations from the API
func commandMap(cfg *Config, url *string) error {	
	locations := Locations{}

	err := pokeapi.GetData(url, cfg.apiClient, cfg.cache, &locations)
	if err != nil {
		return err
	}

	cfg.Previous = locations.Previous
	cfg.Next = locations.Next

	for _, location := range(locations.Results) {
		fmt.Println(location.Name)
	}
	return nil
}

type Locations struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}