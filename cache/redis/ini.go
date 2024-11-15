package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/conf"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client
var KeyList []string

func Init(cfg *conf.RedisConfig) (err error) {

	addr := cfg.Host + fmt.Sprintf(":%d", cfg.Port)
	fmt.Println("redis addr:", addr)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr, // Redis服务器地址
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return
}

// Set 增，改
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// Delete 删
func Delete(key string) error {
	return rdb.Del(ctx, key).Err()
}

// Get 查
func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
