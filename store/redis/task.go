package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wedo-workflow/wedo/model"
)

// TaskCreate create a task
func (r *Redis) TaskCreate(ctx context.Context, task *model.Task) error {
	if err := r.db.HSet(ctx, fmt.Sprintf(processInstanceTasks, task.ProcessInstanceId), task.Id, task).Err(); err != nil {
		return err
	}
	if err := r.db.Set(ctx, fmt.Sprintf(taskDetail, task.Id), task, redis.KeepTTL).Err(); err != nil {
		return err
	}
	return nil
}
