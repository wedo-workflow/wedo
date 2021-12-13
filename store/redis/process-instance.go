package redis

import (
	"context"
	"fmt"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) ProcessInstanceGet(ctx context.Context, processInstanceId string) (*model.ProcessInstance, error) {
	pi := &model.ProcessInstance{}
	if err := r.db.Get(ctx, fmt.Sprintf(processInstanceDetail, processInstanceId)).Scan(pi); err != nil {
		return nil, err
	}
	return pi, nil
}
