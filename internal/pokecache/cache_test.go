package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key	string
		val	[]byte
	} {
		{
			key:	"https://example.com",
			val:	[]byte("testing data"),
		},
		{
			key:	"https://example.com/path",
			val:	[]byte("more data"),
		},
	}

	for i, c := range(cases) {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T){
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("expected to find a key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find a value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond	
	testKey := "https://example.com"
	testVal := []byte("testing data")


	cache := NewCache(baseTime)
	cache.Add(testKey, testVal)

	_, exists := cache.Get(testKey)
	if !exists {
		t.Errorf("expected to find a key")
		return
	}

	time.Sleep(waitTime)

	_, exists = cache.Get(testKey)
	if exists {
		t.Errorf("expected not to find a key")
		return
	}
	
}