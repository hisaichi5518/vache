package vache

import "time"

type CacheData struct {
	Value  string
	Expire time.Duration
}

var cache map[string]CacheData = make(map[string]CacheData)

func Set(key string, val string, expire time.Duration) {
	cache[key] = CacheData{Value: val, Expire: expire}

	time.AfterFunc(expire, func() {
		cache[key] = CacheData{}
	})
}

func Get(key string) string {
	cacheData := cache[key]
	return cacheData.Value
}
