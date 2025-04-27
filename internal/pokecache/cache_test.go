package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache()
	
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("val1"),
		},
		{
			key: "key2",
			val: []byte("val2"),
		},
	}
	
	for _, c := range cases {
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %s in cache", c.key)
			continue
		}
		if string(val) != string(c.val) {
			t.Errorf("expected to find value %s in cache, got %s", string(c.val), string(val))
		}
	}
}

func TestReap(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("val1"))
	
	// Wait for the entry to expire
	time.Sleep(time.Second * 2)
	cache.reap(time.Second)
	
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("expected key1 to be reaped from cache")
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("val1"))
	
	// Start reaping every 1 second
	cache.reapLoop(time.Second)
	
	// Wait for the reaper to run
	time.Sleep(time.Second * 2)
	
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("expected key1 to be reaped from cache")
	}
}
