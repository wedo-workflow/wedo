package runtime

import (
	"context"
	"errors"
	"time"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrInvalidProcessInstanceId = errors.New("invalid process instance id")
)

// TaskCreate creates a new task.
func (r *Runtime) TaskCreate(ctx context.Context, task *wedo.TaskCreateRequest) (string, error) {
	task.Id = r.generateUUID()
	now := time.Now()
	if task.ProcessInstanceId == "" {
		return "", ErrInvalidProcessInstanceId
	}
	newTask := &model.Task{
		CreatedAt:         now,
		UpdatedAt:         now,
		Id:                task.Id,
		Name:              task.Name,
		Description:       task.Description,
		Assignee:          task.Assignee,
		Owner:             task.Owner,
		DelegationState:   task.DelegationState,
		Due:               timestamppb.New(now),
		Priority:          task.Priority,
		ParentTaskId:      task.ParentTaskId,
		ProcessInstanceId: task.ProcessInstanceId,
		TaskDefinitionKey: task.TaskDefinitionKey,
		NamespaceId:       task.NamespaceId,
		FormKey:           task.FormKey,
		FollowUpDate:      timestamppb.New(now),
	}
	if err := r.store.TaskCreate(ctx, newTask); err != nil {
		return "", err
	}
	return task.Id, nil
}
