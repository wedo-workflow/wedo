package redis

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, key, value string) error {
	return r.db.HSet(ctx, processDefinition, key, value).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, key string) (string, error) {
	result := r.db.HGet(ctx, processDefinition, key)
	return result.Val(), result.Err()
}

// ProcessDefinitionStart start a process definition instance
func (r *Redis) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	// var deploy = &model.Deployment{}
	// err := r.db.Get(ctx, fmt.Sprintf(deploySetStart, processID)).Scan(deploy)
	// return deploy, err // processInstanceID
	return "", nil
}
