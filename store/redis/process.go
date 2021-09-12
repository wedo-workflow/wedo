package redis

import (
	"context"
	"fmt"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, deploy *model.Deploy) error {
	return r.db.HSet(ctx, "ProcessDefinitions", deploy.DID, deploy.Name).Err()
}

func (r *Redis) ProcessDefinition(ctx context.Context, processID string) (*model.Deploy, error) {
	var deploy = &model.Deploy{}
	err := r.db.Get(ctx, fmt.Sprintf("%s_deploy", processID)).Scan(deploy)
	return deploy, err
}
