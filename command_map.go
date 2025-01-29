package main

import (
	"fmt"
)

// Prints next 20 locations from the API
func commandMap(cfg *Config) error {
	url := cfg.Next	
	locations := Locations{}

	err := getData(url, &locations)
	if err != nil {
		return err
	}

	switch v := locations.Previous.(type) {
	case string:
		cfg.Previous = v
	default:
		cfg.Previous = ""
	}
	cfg.Next = locations.Next

	for i := 0; i < len(locations.Results); i++{
		fmt.Println(locations.Results[i].Name)
	}
	return nil
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}