package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct{ cache *redis.Client }

var _ Cache = (*Redis)(nil)

func NewRedis(cache *redis.Client) *Redis {
	return &Redis{cache: cache}
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	v, err := r.cache.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return v, nil
}

func (r *Redis) Set(ctx context.Context, key string, value []byte, expiry time.Duration) error {
	_, err := r.cache.Set(ctx, key, value, expiry).Result()
	return err
}

func (r *Redis) Del(ctx context.Context, key string) error {
	_, err := r.cache.Del(ctx, key).Result()
	return err
}
