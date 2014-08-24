package vache

import (
	"bytes"
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	Set("key1", "val1", time.Second)
	Set("key2", "val2", time.Second)

	// found cache.
	if v, ok := Get("key1"); !ok || v != "val1" {
		t.Fatal("not set")
	}

	time.Sleep(1 * time.Second)

	// expired. so, not found cache.
	if _, ok := Get("key1"); ok {
		t.Fatal("not expired!")
	}
	if _, ok := Get("key2"); ok {
		t.Fatal("not expired!")
	}
}

func TestValueIsNotFound(t *testing.T) {
	if v, ok := Get("not-found"); v != "" || ok {
		t.Fatal("found 'not-found-key's cache.")
	}
}

func TestGetOrSet(t *testing.T) {
	key := "get_or_set"
	if v, ok := Get(key); v != "" || ok {
		t.Fatal("found 'get_or_set's cache.")
	}

	value, isGot := GetOrSet(key, func() (string, time.Duration) {
		return "value1", time.Second
	})
	if value != "value1" || isGot {
		t.Fatal("not set")
	}

	value, isGot = GetOrSet(key, func() (string, time.Duration) {
		return "value2", time.Second
	})
	if value == "value2" || !isGot {
		t.Fatal("set value2")
	}

	time.Sleep(1 * time.Second)

	// expired. so, not found cache.
	if value, ok := Get(key); value != "" || ok {
		t.Fatal("not expired")
	}
}

func TestDelete(t *testing.T) {
	value, _ := GetOrSet("name", func() (string, time.Duration) {
		return "hisaichi5518", 10 * time.Second
	})
	if value == "" {
		t.Fatal("not set")
	}

	if deletedValue, ok := Delete("name"); deletedValue == "" || !ok {
		t.Fatal("not delete")
	}
}

func TestLock(t *testing.T) {
	Set("name", "hisaichi5518", 10*time.Second)

	ch := make(chan string)
	go func() {
		Get("name")
		ch <- "1"
	}()
	go func() {
		Get("name")
		ch <- "2"
	}()
	go func() {
		Set("name", "hisaichi5518", 10*time.Second)
		ch <- "3"
	}()
	go func() {
		Get("name")
		ch <- "4"
	}()
	go func() {
		Delete("name")
		ch <- "5"
	}()
	go func() {
		Get("name")
		ch <- "6"
	}()

	var buffer bytes.Buffer
	for i := 0; i < 6; i++ {
		buffer.WriteString(<-ch)
	}
	if buffer.String() != "123456" {
		t.Fatal("not locked", buffer.String())
	}
}
