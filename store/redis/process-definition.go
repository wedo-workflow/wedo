package redis

import (
	"context"
	"fmt"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, deploy *model.Deployment) error {
	return r.db.HSet(ctx, processDefinition, deploy.DID, deploy.Name).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, processID string) (*model.Deployment, error) {
	var deploy = &model.Deployment{}
	err := r.db.Get(ctx, fmt.Sprintf(deploySet, processID)).Scan(deploy)
	return deploy, err
}

// ProcessDefinitionStart start a process definition instance
func (r *Redis) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	// var deploy = &model.Deployment{}
	// err := r.db.Get(ctx, fmt.Sprintf(deploySetStart, processID)).Scan(deploy)
	// return deploy, err // processInstanceID
	return "", nil
}
