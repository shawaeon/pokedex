package main

import (
	"fmt"
)

func commandMapForward(cfg *Config) error {
	url := cfg.Next
	if url == nil {
		fmt.Println("No more locations")
		return nil
	} 
	return commandMap(cfg, url)
}

func commandMapBack(cfg *Config) error {
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

	err := getData(url, cfg.apiClient, cfg.cache, &locations)
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