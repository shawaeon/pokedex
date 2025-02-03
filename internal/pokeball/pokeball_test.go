package pokeball

import (
	"fmt"
	"testing"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key	string
		val	Pokemon
	} {
		{
			key:	"pikachu",
			val:	Pokemon{Name: "pikachu"},
		},
		{
			key:	"charmander",
			val:	Pokemon{Name: "charmander"},
		},
	}
	for i, c := range(cases) {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T){
			pokeball := NewPokeball()
			pokeball.Add(c.key, c.val)
			val, exists := pokeball.Get(c.key)
			if !exists {
				t.Errorf("expected to find a key")
			}
			if val.Name != c.val.Name {
				t.Errorf("expected to find a value")
			}
		})
	}
}