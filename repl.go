package main

import "strings"

// Convert to lowercase and split based on whitespaces
func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}