package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wedo-workflow/wedo/element/wedo_model"
)

func (r *Redis) Deploy(ctx context.Context, deployID string) (*wedo_model.Deploy, error) {
	deploy := &wedo_model.Deploy{}
	err := r.db.Get(ctx, deployID).Scan(deploy)
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (r *Redis) DeploySet(ctx context.Context, deploy *wedo_model.Deploy) error {
	if deploy.Name == "" {
		return errors.New("deploy name is empty")
	}
	deployBytes, err := json.Marshal(deploy)
	if err != nil {
		return err
	}
	return r.db.Set(ctx, deploy.Name, deployBytes, 0).Err()
}
