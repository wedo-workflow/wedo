package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/model"
)

func deployKey(deployID string) string {
	return fmt.Sprintf(deploySet, deployID)
}

func (r *Redis) Deploy(ctx context.Context, deployID string) (*model.Deploy, error) {
	deploy := &model.Deploy{}
	result, err := r.db.Get(ctx, deployKey(deployID)).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(result), deploy); err != nil {
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
	return r.db.Set(ctx, deployKey(deploy.DID), deployBytes, 0).Err()
}

// DeployList returns a list of deployments.
func (r *Redis) DeployList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deploy, error) {
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
		deployKeys = append(deployKeys, deployKey(key))
	}
	deploys := make([]*model.Deploy, 0)
	results, err := r.db.MGet(ctx, deployKeys...).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		if result == nil {
			log.Debug("deployment list result is nil")
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

// DeployDelete deletes a deployment.
func (r *Redis) DeployDelete(ctx context.Context, deployID string) error {
	return r.db.Del(ctx, deployKey(deployID)).Err()
}
