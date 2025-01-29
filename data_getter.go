package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Gets data from the API and adds it to a struct
func getData[T any](url string, data *T) error {
	
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making a request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}
	
	
	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}
	return nil
}
