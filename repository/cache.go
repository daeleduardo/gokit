package database

import (
	"net/url"
	"github.com/go-redis/redis/v8"
	"gokit/config"
)


//Return a Redis Connection
func RedisConn() *redis.Client {
	Url, _ := url.Parse(config.GetConfig().CacheConn)
	Pass, _ := Url.User.Password()
	return redis.NewClient(&redis.Options{
		Addr:     Url.Host,
		Password: Pass,
	})
}
