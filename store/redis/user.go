package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) UserCreate(ctx context.Context, user *model.User) error {
	now := time.Now().UnixMilli()
	if err := r.db.ZAdd(ctx, userList, &redis.Z{Score: float64(now), Member: user.Id}).Err(); err != nil {
		return err
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.db.Set(ctx, fmt.Sprintf(userProfile, user.Id), string(userBytes), redis.KeepTTL).Err()
}

// UserGet ByID returns the user with the given ID.
func (r *Redis) UserGet(ctx context.Context, id string) (*model.User, error) {
	userBytes, err := r.db.Get(ctx, fmt.Sprintf(userProfile, id)).Bytes()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserDelete deletes the user with the given ID.
func (r *Redis) UserDelete(ctx context.Context, id string) error {
	if err := r.db.ZRem(ctx, userList, id).Err(); err != nil {
		return err
	}
	return r.db.Del(ctx, fmt.Sprintf(userProfile, id)).Err()
}
