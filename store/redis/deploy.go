package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) Deploy(ctx context.Context, deployID string) (*model.Deploy, error) {
	deploy := &model.Deploy{}
	err := r.db.Get(ctx, deployID).Scan(deploy)
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (r *Redis) DeploySet(ctx context.Context, deploy *model.Deploy) error {
	if deploy.DID == "" {
		return errors.New("deploy id is empty")
	}
	if deploy.Name == "" {
		return errors.New("deploy name is empty")
	}
	deploy.CreateTime = time.Now()
	deployBytes, err := json.Marshal(deploy)
	if err != nil {
		return err
	}
	return r.db.Set(ctx, fmt.Sprintf(deploySet, deploy.DID), deployBytes, 0).Err()
}

// DeploymentList returns a list of deployments.
func (r *Redis) DeploymentList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deploy, error) {
	type DeployKey struct {
		Key string
		Val string
	}
	deployKeys := []string{}
	lists, err := r.db.HGetAll(ctx, processDefinition).Result()
	if err != nil {
		return nil, err
	}
	for key, _ := range lists {
		deployKeys = append(deployKeys, fmt.Sprintf(deploySet, key))

	}
	deploys := make([]*model.Deploy, 0)
	results, err := r.db.MGet(ctx, deployKeys...).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if result == nil {
			continue
		}
		deploy := &model.Deploy{}
		err := json.Unmarshal([]byte(result.(string)), deploy)
		if err != nil {
			return nil, err
		}
		deploys = append(deploys, deploy)
	}
	if err != nil {
		return nil, err
	}
	return deploys, nil
}
