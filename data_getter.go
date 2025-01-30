package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

// Gets data from the API and adds it to a struct
func getData[T any](url *string, client *http.Client, cache *pokecache.Cache, dataStruct *T) error {
	
	// If data exists in cache use cached data
	data, exists := cache.Get(*url)
	if exists {
		err := json.Unmarshal(data, dataStruct)
		if err != nil {
			return fmt.Errorf("error unmarshalling data: %w", err)
		}
		return nil
	}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return fmt.Errorf("error creating a request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making a request: %w", err)
	}
	defer res.Body.Close()
	
	// Cache data
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}
	cache.Add(*url, data)
	
	err = json.Unmarshal(data, dataStruct)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}
	return nil
}
