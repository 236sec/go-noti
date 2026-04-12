package redis

import (
	"context"
	"time"

	"goboilerplate.com/config"
)

type IRedis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (int64, error)
	HSet(ctx context.Context, key string, values ...interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key, field string) (bool, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
}

func GetRedisClient() IRedis {
	cfg := config.GetConfig()
	return NewRedisClient(&cfg.EnvConfig.Redis)
}
