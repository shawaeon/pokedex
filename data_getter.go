package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Gets data from the API and adds it to a struct
func getData[T any](url *string, client *http.Client, dataStruct *T) error {
	
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return fmt.Errorf("error creating a request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making a request: %w", err)
	}
	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data: %w", err)
	}
	
	err = json.Unmarshal(data, dataStruct)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}
	return nil
}
