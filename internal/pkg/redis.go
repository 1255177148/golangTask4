// 定义公共变量，这样防止依赖循环

package pkg

import "github.com/redis/go-redis/v9"

var RDB *redis.Client
