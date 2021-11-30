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

// TaskGet gets a task by id.
func (r *Runtime) TaskGet(ctx context.Context, taskId string) (*wedo.TaskResponse, error) {
	task, err := r.store.TaskGet(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return &wedo.TaskResponse{
		Id:                task.Id,
		Name:              task.Name,
		Description:       task.Description,
		Assignee:          task.Assignee,
		Owner:             task.Owner,
		DelegationState:   task.DelegationState,
		Due:               task.Due,
		Priority:          task.Priority,
		ParentTaskId:      task.ParentTaskId,
		ProcessInstanceId: task.ProcessInstanceId,
		TaskDefinitionKey: task.TaskDefinitionKey,
		NamespaceId:       task.NamespaceId,
		FormKey:           task.FormKey,
		FollowUpDate:      task.FollowUpDate,
	}, nil
}

// TaskDelete deletes a task by id.
func (r *Runtime) TaskDelete(ctx context.Context, taskId string) error {
	oldTaks, err := r.store.TaskGet(ctx, taskId)
	if err != nil {
		return err
	}
	return r.store.TaskDelete(ctx, taskId, oldTaks.ProcessInstanceId)
}

// TaskList gets a list of tasks by process instance id.
func (r *Runtime) TaskList(ctx context.Context, request *wedo.TaskListRequest) ([]*wedo.TaskResponse, error) {
	tasks, err := r.store.TaskList(ctx, &model.TaskListOptions{ProcessInstanceID: request.ProcessInstanceId})
	if err != nil {
		return nil, err
	}
	response := make([]*wedo.TaskResponse, len(tasks))
	for i, task := range tasks {
		response[i] = &wedo.TaskResponse{
			Id:                task.Id,
			Name:              task.Name,
			Description:       task.Description,
			Assignee:          task.Assignee,
			Owner:             task.Owner,
			DelegationState:   task.DelegationState,
			Due:               task.Due,
			Priority:          task.Priority,
			ParentTaskId:      task.ParentTaskId,
			ProcessInstanceId: task.ProcessInstanceId,
			TaskDefinitionKey: task.TaskDefinitionKey,
			NamespaceId:       task.NamespaceId,
			FormKey:           task.FormKey,
			FollowUpDate:      task.FollowUpDate,
		}
	}
	return response, nil
}
