package vache

import (
	"sync"
	"time"
)

var mu sync.RWMutex = sync.RWMutex{}
var cache map[string]string = make(map[string]string)

func Set(key, val string, expire time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	cache[key] = val

	time.AfterFunc(expire, func() {
		mu.Lock()
		defer mu.Unlock()
		delete(cache, key)
	})
}

func Get(key string) (value string, ok bool) {
	mu.RLock()
	defer mu.RUnlock()

	value, ok = cache[key]
	return
}

func GetOrSet(key string, code func() (string, time.Duration)) (string, bool) {
	if value, ok := Get(key); ok {
		return value, true
	}

	value, expire := code()
	Set(key, value, expire)
	return value, false
}

func Delete(key string) (deletedValue string, ok bool) {
	mu.Lock()
	defer mu.Unlock()

	deletedValue, ok = cache[key]
	delete(cache, key)
	return
}
