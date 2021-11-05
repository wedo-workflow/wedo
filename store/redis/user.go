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
