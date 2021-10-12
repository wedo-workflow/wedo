package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
)

func (APIServer) UserCreate(ctx context.Context, request *wedo.UserCreateRequest) (*wedo.UserCreateResponse, error) {
	panic("implement me")
}

func (APIServer) UserGet(ctx context.Context, request *wedo.UserRequest) (*wedo.UserResponse, error) {
	panic("implement me")
}

func (APIServer) UserDelete(ctx context.Context, request *wedo.UserDeleteRequest) (*wedo.UserDeleteResponse, error) {
	panic("implement me")
}
