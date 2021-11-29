package redis

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/model"
)

func deployKey(deployID string) string {
	return fmt.Sprintf(deploySet, deployID)
}

func (r *Redis) Deployment(ctx context.Context, deployID string) (*model.Deployment, error) {
	deploy := &model.Deployment{}
	result, err := r.db.Get(ctx, deployKey(deployID)).Result()
	if err != nil {
		return nil, err
	}
	if err := deploy.UnmarshalBinary([]byte(result)); err != nil {
		return nil, err
	}
	return deploy, nil
}

func (r *Redis) DeploymentCreate(ctx context.Context, deploy *model.Deployment) error {
	return r.db.HSet(ctx, deploys, deploy.DID, deploy).Err()
}

// DeploymentList returns a list of deployments.
func (r *Redis) DeploymentList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deployment, error) {
	type DeployKey struct {
		Key string
		Val string
	}
	deployKeys := []string{}
	lists, err := r.db.HGetAll(ctx, deploys).Result()
	if err != nil {
		return nil, err
	}
	for key, _ := range lists {
		deployKeys = append(deployKeys, deployKey(key))
	}
	deploys := make([]*model.Deployment, 0)
	if len(deployKeys) == 0 {
		return deploys, nil
	}
	results, err := r.db.MGet(ctx, deployKeys...).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if result == nil {
			log.Debug("deployment list result is nil")
			continue
		}
		deploy := &model.Deployment{}
		if err := deploy.UnmarshalBinary([]byte(result.(string))); err != nil {
			return nil, err
		}
		deploys = append(deploys, deploy)
	}
	if err != nil {
		return nil, err
	}
	return deploys, nil
}

// DeployDelete deletes a deployment.
func (r *Redis) DeploymentDelete(ctx context.Context, deployID string) error {
	return r.db.Del(ctx, deployKey(deployID)).Err()
}
