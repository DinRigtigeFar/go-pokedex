package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))

	if val, ok := cache.Get("key1"); !ok || string(val) != "value1" {
		t.Errorf("Expected value1 for key1, got %v", val)
	}

	if val, ok := cache.Get("key2"); !ok || string(val) != "value2" {
		t.Errorf("Expected value2 for key2, got %v", val)
	}

	if _, ok := cache.Get("key3"); ok {
		t.Error("Expected no value for key3, but got one")
	}
	
	time.Sleep(6 * waitTime)

	if _, ok := cache.Get("key1"); ok {
		t.Error("Expected cache to be cleared, but got a value for key1")
	}

	if _, ok := cache.Get("key2"); ok{
		t.Error("Expected cache to be cleared, but got a value for key2")
	}

	if _, ok := cache.Get("key3"); ok {
		t.Error("Expected cache to be cleared and no value for key3, but got a value for it")
	}
}