package main

import "fmt"

func commandExplore(cfg *Config) error {
	if cfg.optionalParam == "" {
		fmt.Println("Please input location to explore (f.ex. \"explore canalave-city-area\").")
		return nil
	}
	fmt.Println(cfg.optionalParam)
	return nil
}