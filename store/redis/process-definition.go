package redis

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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
	rest := r.db.HSet(ctx, processInstance, pd.Id)
	if rest.Err() != nil {
		return "", rest.Err()
	}

	piIDV4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	piID := piIDV4.String()

	pInstance := &model.ProcessIntance{
		Status: model.PiOK,
	}
	if err := r.db.HSet(ctx, fmt.Sprintf(processInstanceDetail, piID), pInstance).Err(); err != nil {
		return "", err
	}
	return piID, nil
}
