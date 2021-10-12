package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
)

func (APIServer) TaskCreate(ctx context.Context, request *wedo.TaskCreateRequest) (*wedo.TaskCreateResponse, error) {
	panic("implement me")
}

func (APIServer) TaskGet(ctx context.Context, request *wedo.TaskRequest) (*wedo.TaskResponse, error) {
	panic("implement me")
}

func (APIServer) TaskDelete(ctx context.Context, request *wedo.TaskDeleteRequest) (*wedo.TaskDeleteResponse, error) {
	panic("implement me")
}
