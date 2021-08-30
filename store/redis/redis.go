package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/element"
	"github.com/wedo-workflow/wedo/runtime/config"
)

type Redis struct {
	db *redis.Client
}

func NewRedis(config *config.Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Store.DSN,
		Password: config.Store.Password,
		DB:       0, // use default DB
	})
	return &Redis{
		db: rdb,
	}, nil
}

func (r *Redis) ElementSave(ctx context.Context, element element.Element, rootID string) error {
	if element == nil {
		return errors.New("element is nil")
	}
	if rootID == "" {
		return errors.New("rootID is empty")
	}
	elementBytes, err := json.Marshal(element)
	if err != nil {
		return err
	}
	if err := r.db.HSet(ctx, rootID, element.EID(), elementBytes).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.db.Ping(ctx).Err()
}
