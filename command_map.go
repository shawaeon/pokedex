package main

import (
	"fmt"
)

func commandMapBack(cfg *Config) error {
	url := cfg.Previous	
	if url == nil {
		fmt.Println("You're on the first page")
		return nil
	} 

	locations := Locations{}

	err := getData(url, &locations)
	if err != nil {
		return err
	}
	
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next

	for i := 0; i < len(locations.Results); i++{
		fmt.Println(locations.Results[i].Name)
	}
	return nil
}

// Prints next 20 locations from the API
func commandMap(cfg *Config) error {
	url := cfg.Next
	if url == nil {
		fmt.Println("No more locations")
		return nil
	} 
	
	locations := Locations{}

	err := getData(url, &locations)
	if err != nil {
		return err
	}

	cfg.Previous = locations.Previous
	cfg.Next = locations.Next

	for i := 0; i < len(locations.Results); i++{
		fmt.Println(locations.Results[i].Name)
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