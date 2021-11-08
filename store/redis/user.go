package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) UserCreate(ctx context.Context, user *model.User) error {
	now := time.Now().UnixMilli()
	if err := r.db.ZAdd(ctx, userList, &redis.Z{Score: float64(now), Member: user.Id}).Err(); err != nil {
		return err
	}
	return r.db.Set(ctx, fmt.Sprintf(userProfile, user.Id), user, redis.KeepTTL).Err()
}

// UserGet ByID returns the user with the given ID.
func (r *Redis) UserGet(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.Get(ctx, fmt.Sprintf(userProfile, id)).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UserDelete deletes the user with the given ID.
func (r *Redis) UserDelete(ctx context.Context, id string) error {
	if err := r.db.ZRem(ctx, userList, id).Err(); err != nil {
		return err
	}
	return r.db.Del(ctx, fmt.Sprintf(userProfile, id)).Err()
}

// UserList returns a list of all users.
func (r *Redis) UserList(ctx context.Context, opts *model.UserListOptions) ([]*model.User, error) {
	userIDs := make([]string, 0)
	if err := r.db.ZRevRange(ctx, userList, opts.Offset, opts.Offset+opts.Limit-1).ScanSlice(&userIDs); err != nil {
		return nil, err
	}
	userIDKeys := make([]string, len(userIDs))
	for _, id := range userIDs {
		userIDKeys = append(userIDKeys, fmt.Sprintf(userProfile, id))
	}
	users := make([]*model.User, 0, len(userIDs))
	results, err := r.db.MGet(ctx, userIDKeys...).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if result == nil {
			log.Debug("deployment list result is nil")
			continue
		}
		user := &model.User{}
		if err := user.UnmarshalBinary([]byte(result.(string))); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err != nil {
		return nil, err
	}

	return users, nil
}
