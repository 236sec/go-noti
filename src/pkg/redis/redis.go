package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"goboilerplate.com/config"
)

type Redis struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.RedisConfig) IRedis {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.RedisPassword,
		DB:           cfg.RedisDB,
		Protocol:     cfg.Protocol,
	})

	return &Redis{client: client}
}

func (w *Redis) Get(ctx context.Context, key string) (string, error) {
	return w.client.Get(ctx, key).Result()
}

func (w *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return w.client.Set(ctx, key, value, expiration).Err()
}

func (w *Redis) Del(ctx context.Context, keys ...string) error {
	return w.client.Del(ctx, keys...).Err()
}

func (w *Redis) Exists(ctx context.Context, keys ...string) (int64, error) {
	return w.client.Exists(ctx, keys...).Result()
}

func (w *Redis) HSet(ctx context.Context, key string, values ...interface{}) error {
	return w.client.HSet(ctx, key, values...).Err()
}

func (w *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	return w.client.HGet(ctx, key, field).Result()
}

func (w *Redis) HDel(ctx context.Context, key string, fields ...string) error {
	return w.client.HDel(ctx, key, fields...).Err()
}

func (w *Redis) HExists(ctx context.Context, key, field string) (bool, error) {
	return w.client.HExists(ctx, key, field).Result()
}

func (w *Redis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return w.client.HGetAll(ctx, key).Result()
}
