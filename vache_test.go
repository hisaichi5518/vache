package vache

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	Set("key1", "val1", time.Second)
	Set("key2", "val2", time.Second)

	// found cache.
	v := Get("key1")
	if v != "val1" {
		t.Fatal("not set")
	}
	time.Sleep(2 * time.Second)

	// expired. so, not found cache.
	v = Get("key1")
	if v != "" {
		t.Fatal("not expired!")
	}
}

func TestValueIsNotFound(t *testing.T) {
	v := Get("not-found")
	if v != "" {
		t.Fatal("found 'not-found-key's cache.")
	}
}

func TestGetOrSet(t *testing.T) {
	key := "get_or_set"
	v := Get(key)
	if v != "" {
		t.Fatal("found 'get_or_set's cache.")
	}

	value := GetOrSet(key, func() (string, time.Duration) {
		return "value1", time.Second
	})
	if value != "value1" {
		t.Fatal("not set")
	}

	value = GetOrSet(key, func() (string, time.Duration) {
		return "value2", time.Second
	})
	if value == "value2" {
		t.Fatal("set value2")
	}

	time.Sleep(2 * time.Second)

	// expired. so, not found cache.
	value = Get(key)
	if value != "" {
		t.Fatal("not expired")
	}
}
