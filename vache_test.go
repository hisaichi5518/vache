package vache

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	Set("key1", "val1", time.Second)
	Set("key2", "val2", time.Second)

	// found cache.
	v := Get("key1")
	if v != "val1" {
		t.Fatal(v + " is not val1")
	}
	time.Sleep(2 * time.Second)

	// expired. so, not found cache.
	v = Get("key1")
	if v != "" {
		t.Fatal(" is not val1")
	}
}

func TestValueIsNotFound(t *testing.T) {
	v := Get("not-found")
	if v != "" {
		t.Fatal("found 'not-found-key's cache.")
	}
}
