package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessInstanceGet(ctx context.Context, processInstanceId string) (*model.ProcessInstance, error) {
	pi := &model.ProcessInstance{}
	if err := r.db.Get(ctx, fmt.Sprintf(processInstanceDetail, processInstanceId)).Scan(pi); err != nil {
		return nil, err
	}
	return pi, nil
}

func (r *Redis) ProcessInstanceUpdateStatus(ctx context.Context, pi *model.ProcessInstance) error {
	return r.db.Set(ctx, fmt.Sprintf(processInstanceStatus, pi.Id), pi.Status.String(), redis.KeepTTL).Err()
}
