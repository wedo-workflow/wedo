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

// TaskGet get a task
func (r *Redis) TaskGet(ctx context.Context, taskId string) (*model.Task, error) {
	var task model.Task
	if err := r.db.Get(ctx, fmt.Sprintf(taskDetail, taskId)).Scan(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

// TaskDelete delete a task
func (r *Redis) TaskDelete(ctx context.Context, taskID string, processInstanceID string) error {
	if err := r.db.Del(ctx, fmt.Sprintf(taskDetail, taskID)).Err(); err != nil {
		return err
	}
	return r.db.HDel(ctx, fmt.Sprintf(processInstanceTasks, processInstanceID), taskID).Err()
}

func (r *Redis) TaskList(ctx context.Context, opts *model.TaskListOptions) ([]*model.Task, error) {
	var tasks []*model.Task
	results, err := r.db.HGetAll(ctx, fmt.Sprintf(processInstanceTasks, opts.ProcessInstanceID)).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		task := &model.Task{}
		if err := task.UnmarshalBinary([]byte(result)); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
