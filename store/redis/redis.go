package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/runtime/config"
)

type Redis struct {
	db *redis.Client
}

func NewRedis(config *config.Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Store.Redis.DSN,
		Password: config.Store.Redis.Password,
		DB:       config.Store.Redis.DB,
	})
	return &Redis{
		db: rdb,
	}, nil
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.db.Ping(ctx).Err()
}
