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

func (APIServer) TaskGet(ctx context.Context, request *wedo.TaskRequest) (*wedo.TaskResponse, error) {
	panic("implement me")
}

func (APIServer) TaskDelete(ctx context.Context, request *wedo.TaskDeleteRequest) (*wedo.TaskDeleteResponse, error) {
	panic("implement me")
}

// TaskList lists all tasks.
func (APIServer) TaskList(ctx context.Context, request *wedo.TaskListRequest) (*wedo.TaskListResponse, error) {
	panic("implement me")
}
