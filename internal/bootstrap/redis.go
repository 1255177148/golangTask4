package bootstrap

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: config.Cfg.Redis.Host,
	})
}
