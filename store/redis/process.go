package redis

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessDefinitionAdd(ctx context.Context, deploy *model.Deploy) error {
	return r.db.HSet(ctx, "ProcessDefinitions", deploy.DID, deploy.Name).Err()
}
