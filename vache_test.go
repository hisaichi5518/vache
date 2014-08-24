package vache

import (
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	Set("key1", "val1", time.Second)
	Set("key2", "val2", time.Second)

	// found cache.
	v, _ := Get("key1")
	if v != "val1" {
		t.Fatal("not set")
	}
	time.Sleep(2 * time.Second)

	// expired. so, not found cache.
	v, _ = Get("key1")
	if v != "" {
		t.Fatal("not expired!")
	}
}

func TestValueIsNotFound(t *testing.T) {
	v, _ := Get("not-found")
	if v != "" {
		t.Fatal("found 'not-found-key's cache.")
	}
}

func TestGetOrSet(t *testing.T) {
	key := "get_or_set"
	v, _ := Get(key)
	if v != "" {
		t.Fatal("found 'get_or_set's cache.")
	}

	value, _ := GetOrSet(key, func() (string, time.Duration) {
		return "value1", time.Second
	})
	if value != "value1" {
		t.Fatal("not set")
	}

	value, _ = GetOrSet(key, func() (string, time.Duration) {
		return "value2", time.Second
	})
	if value == "value2" {
		t.Fatal("set value2")
	}

	time.Sleep(2 * time.Second)

	// expired. so, not found cache.
	value, _ = Get(key)
	if value != "" {
		t.Fatal("not expired")
	}
}

func TestDelete(t *testing.T) {
	Set("name", "hisaichi5518", 10*time.Second)
	if v, _ := Get("name"); v == "" {
		t.Fatal("not set")
	}

	Delete("name")
	if v, _ := Get("name"); v != "" {
		t.Fatal("not delete")
	}
}

func TestLock(t *testing.T) {
	Set("name", "hisaichi5518", 10*time.Second)

	go func() {
		v, _ := Get("name")
		if v != "hisaichi5518" {
			t.Fatal("not get")
		}
	}()
	go func() {
		v, _ := Get("name")
		if v != "hisaichi5518" {
			t.Fatal("not get")
		}
	}()
	go Delete("name")
	go Set("name", "hisaichi5518", 10*time.Second)
	go func() {
		v, _ := Get("name")
		if v != "hisaichi5518" {
			t.Fatal("not get")
		}
	}()
	time.Sleep(10 * time.Second)
}
