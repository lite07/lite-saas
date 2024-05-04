package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func SetCacheWithExpiry(key string, value interface{}, expiration time.Duration) {
	getCacheInstance().Set(key, value, expiration)
}

func SetCache(key string, value interface{}) {
	getCacheInstance().Set(key, value, cache.DefaultExpiration)
}

func GetCache(key string) any {
	if value, found := getCacheInstance().Get(key); found {
		return value
	} else {
		return nil
	}
}

var cacheInstance *cache.Cache

func getCacheInstance() *cache.Cache {
	if cacheInstance == nil {
		cacheInstance = cache.New(1*time.Minute, 5*time.Minute)
	}

	return cacheInstance
}
