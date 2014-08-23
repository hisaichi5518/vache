package vache

import "time"

var cache map[string]string = make(map[string]string)

func Set(key string, val string, expire time.Duration) {
	cache[key] = val

	time.AfterFunc(expire, func() {
		cache[key] = ""
	})
}

func Get(key string) string {
	cacheData := cache[key]
	return cacheData
}
