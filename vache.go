package vache

import "time"

var cache map[string]string = make(map[string]string)

func Set(key string, val string, expire time.Duration) {
	cache[key] = val

	time.AfterFunc(expire, func() {
		delete(cache, key)
	})
}

func Get(key string) string {
	cacheData := cache[key]
	return cacheData
}

func GetOrSet(key string, code func() (string, time.Duration)) string {
	v := Get(key)
	if v != "" {
		return v
	}

	var expire time.Duration
	v, expire = code()
	Set(key, v, expire)
	return v
}

func Delete(key string) string {
	v := Get(key)
	delete(cache, key)

	return v
}
