package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/runtime/config"
)

// Redis Keys
var (
	processDefinition = "process_definitions"

	deploys    = "deploys"
	deploySet  = "%s_deploy"
	elementSet = "%s_elements"

	Namespaces    = "namespaces"
	NamespacesSet = "namespaces_set"

	userProfile     = "%s_user"
	userList        = "user_list"
	userListByEmail = "user_list_by_email"
)

type Redis struct {
	db *redis.Client
}

func NewRedis(config *config.Config) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Store.Redis.Addr, config.Store.Redis.Port),
		Password: config.Store.Redis.Password,
		DB:       config.Store.Redis.DBID,
	})
	return &Redis{
		db: rdb,
	}, nil
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.db.Ping(ctx).Err()
}
