package main

import "fmt"

func commandInspect (cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please input pokemon to inspect (f.ex. \"inspect pikachu\")")
	}
	fmt.Printf("%v\n", args[0])
	return nil
}