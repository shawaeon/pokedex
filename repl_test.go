package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},{
			input: "",
			expected: []string{},
		},{
			input: "single",
			expected: []string{"single"},
		},

	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf(
				"TestCleanInput() length = %v,expected length = %v", 
				len(actual), 
				len(c.expected),
			)
		}
		for i := range(actual) {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(
					"TestCleanInput() string at index %d = %v, expected string = %v", 
					i, 
					word, 
					expectedWord,
				)
			}
		}
	}
}