package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	wedo "github.com/wedo-workflow/wedo/api"

	"github.com/google/uuid"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, pd *model.ProcessDefinition) error {
	if err := r.db.HSet(ctx, fmt.Sprintf(processDefinitionDetail, pd.Id), pd.ToMapping()).Err(); err != nil {
		return err
	}
	return r.db.HSet(ctx, processDefinition, pd.Id, pd).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, key string) (*model.ProcessDefinition, error) {
	var pd *model.ProcessDefinition
	err := r.db.HGet(ctx, processDefinition, key).Scan(&pd)
	return pd, err
}

// ProcessDefinitionStart start a process definition instance
func (r *Redis) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	rest := r.db.HSet(ctx, processInstance, pd.Id)
	if rest.Err() != nil {
		return "", rest.Err()
	}

	piIDV4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	piID := piIDV4.String()

	pInstance := &model.ProcessInstance{
		Status: wedo.ProcessInstanceState_STARTED,
	}
	if err := r.db.Set(ctx, fmt.Sprintf(processInstanceDetail, piID), pInstance, redis.KeepTTL).Err(); err != nil {
		return "", err
	}
	return piID, nil
}
