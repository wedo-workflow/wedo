package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, pd *model.ProcessDefinition) error {
	if err := r.db.HSet(ctx, fmt.Sprintf(processDefinitionDetail, pd.Id), pd.ToMapping()).Err(); err != nil {
		return err
	}
	return r.db.HSet(ctx, processDefinitions, pd.Id, pd).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, key string) (*model.ProcessDefinition, error) {
	pd := &model.ProcessDefinition{}
	err := r.db.HGet(ctx, processDefinitions, key).Scan(pd)
	return pd, err
}

// ProcessDefinitionStart start a process definition instance
func (r *Redis) ProcessDefinitionStart(ctx context.Context, pi *model.ProcessInstance) error {
	if err := r.db.HSet(ctx, fmt.Sprintf(processDefinitionInstances, pi.ProcessDefinition.Id), pi.Id, pi).Err(); err != nil {
		return err
	}
	if err := r.db.HSet(ctx, processInstances, pi.Id, pi).Err(); err != nil {
		return err
	}
	if err := r.db.Set(ctx, fmt.Sprintf(processInstanceDetail, pi.Id), pi, redis.KeepTTL).Err(); err != nil {
		return err
	}
	return nil
}
