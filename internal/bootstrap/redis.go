package bootstrap

import (
	"github.com/1255177148/golangTask4/config"
	"github.com/1255177148/golangTask4/internal/pkg"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	pkg.RDB = redis.NewClient(&redis.Options{
		Addr: config.Cfg.Redis.Host,
	})
}
