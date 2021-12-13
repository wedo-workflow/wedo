package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/runtime/config"
)

// Redis Keys
var (
	processDefinitions         = "process_definitions"
	processDefinitionDetail    = "process_definition_%s_detail"
	processDefinitionInstances = "process_definition_%s_instances"
	processInstances           = "process_instances"
	processInstanceDetail      = "process_instance_%s_detail"
	processInstanceStatus      = "process_instance_%s_status"
	processInstanceTasks       = "process_instance_%s_tasks"

	taskDetail = "tasks_detail_%s"

	deploys    = "deploys"
	deploySet  = "deploys_detail_%s"
	elementSet = "deploys_detail_%s_elements"

	NamespacesByID          = "namespaces_by_id"
	NamespacesByName        = "namespaces_by_name"
	NamespacesNameUniqueSet = "namespaces_name_unqiue_set"

	users              = "users"
	userProfile        = "user_detail_%s"
	userEmailUniqueSet = "users_email_unique_set"
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
