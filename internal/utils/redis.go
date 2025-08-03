package utils

import (
	"context"
	"github.com/1255177148/golangTask4/internal/pkg"
	"time"
)

var ctx = context.Background()

// SetRDB 设置值，支持过期时间
func SetRDB(key string, value interface{}, expiration time.Duration) error {
	return pkg.RDB.Set(ctx, key, value, expiration).Err()
}

// GetRDB 获取值
func GetRDB(key string) (string, error) {
	return pkg.RDB.Get(ctx, key).Result()
}

// ExistsRDB 判断 key 是否存在
func ExistsRDB(key string) (bool, error) {
	val, err := pkg.RDB.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

// DeleteRDB 删除 key
func DeleteRDB(key string) error {
	return pkg.RDB.Del(ctx, key).Err()
}

// DeleteByPrefixRDB 删除匹配前缀的 key
func DeleteByPrefixRDB(prefix string) error {
	iter := pkg.RDB.Scan(ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		pkg.RDB.Del(ctx, iter.Val())
	}
	return iter.Err()
}
