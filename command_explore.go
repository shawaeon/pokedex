package main

import (
	"fmt"
	"strings"

	"github.com/shawaeon/pokedex/internal/pokeapi"
)

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please input location to explore (f.ex. \"explore canalave-city-area\")")
	}
	url := fmt.Sprintf("%v/location-area/%s", pokeapi.BaseURL, args[0])
	location := Location{}

	err:= pokeapi.GetData(&url, cfg.apiClient, cfg.cache, &location)
	if err != nil {
		if strings.Contains(err.Error(), "unmarshalling") {
			return fmt.Errorf("invalid location")
		}
		return err
	}	
	fmt.Printf("Exploring %v\n", location.Location.Name)
	fmt.Printf("Pokemon found in this area:\n")
	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found in area.")
	}
	for _, encounter := range(location.PokemonEncounters) {		
		fmt.Printf(" -%v\n",encounter.Pokemon.Name)
	}
	return nil
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}