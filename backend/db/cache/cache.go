package cache

import (
	"fmt"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"vue-go-users.com/config"
)

var (
	cacheConnection *cache.Cache
)

func Init() {
	conf := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.GetString("REDIS_HOST"), conf.GetString("REDIS_PORT")),
	})

	cacheConnection = cache.New(&cache.Options{
		Redis: rdb,
	})
}

// GetCache returns
func GetCache() *cache.Cache {
	return cacheConnection
}
