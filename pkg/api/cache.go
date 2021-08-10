package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	cacheFilename = "pokeapi-cache.txt"
	cacheDuration = time.Hour * 24 * 7 // 1 week
)

var c *cache.Cache

func setCache(k string, x interface{}) {
	c.Set(k, x, cache.DefaultExpiration)
}

func InitCache() {
	c = cache.New(cacheDuration, cacheDuration)
	c.LoadFile(cacheFilename)
}

func SaveCacheFile() {
	c.SaveFile(cacheFilename)
}
