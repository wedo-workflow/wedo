package redis

import (
	"context"
	"fmt"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, deploy *model.Deploy) error {
	return r.db.HSet(ctx, processDefinition, deploy.DID, deploy.Name).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, processID string) (*model.Deploy, error) {
	var deploy = &model.Deploy{}
	err := r.db.Get(ctx, fmt.Sprintf(deploySet, processID)).Scan(deploy)
	return deploy, err
}

// ProcessDefinitionStart start a process definition instance
func (r *Redis) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	// var deploy = &model.Deploy{}
	// err := r.db.Get(ctx, fmt.Sprintf(deploySetStart, processID)).Scan(deploy)
	// return deploy, err // processInstanceID
	return "", nil
}
