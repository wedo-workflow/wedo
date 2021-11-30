package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) TaskCreate(ctx context.Context, request *wedo.TaskCreateRequest) (*wedo.TaskCreateResponse, error) {
	id, err := s.Runtime.TaskCreate(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskCreateResponse{
		Id:   id,
		Name: request.Name,
	}, nil
}

func (s *APIServer) TaskGet(ctx context.Context, request *wedo.TaskRequest) (*wedo.TaskResponse, error) {
	if request.TaskId == "" {
		return nil, status.Error(codes.InvalidArgument, "task id is empty")
	}
	task, err := s.Runtime.TaskGet(ctx, request.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return task, nil
}

func (s *APIServer) TaskDelete(ctx context.Context, request *wedo.TaskDeleteRequest) (*wedo.TaskDeleteResponse, error) {
	if request.TaskId == "" {
		return nil, status.Error(codes.InvalidArgument, "task id is empty")
	}
	err := s.Runtime.TaskDelete(ctx, request.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskDeleteResponse{}, nil
}

// TaskList lists all tasks.
func (s *APIServer) TaskList(ctx context.Context, request *wedo.TaskListRequest) (*wedo.TaskListResponse, error) {
	if request.ProcessInstanceId == "" {
		return nil, status.Error(codes.InvalidArgument, "process instance id is empty")
	}
	tasks, err := s.Runtime.TaskList(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskListResponse{
		Tasks: tasks,
	}, nil
}
