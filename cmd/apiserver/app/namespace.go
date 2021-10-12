package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
)

func (APIServer) NamespaceCreate(ctx context.Context, request *wedo.NamespaceCreateRequest) (*wedo.NamespaceCreateResponse, error) {
	panic("implement me")
}

func (APIServer) NamespaceGet(ctx context.Context, request *wedo.NamespaceRequest) (*wedo.NamespaceResponse, error) {
	panic("implement me")
}

func (APIServer) NamespaceDelete(ctx context.Context, request *wedo.NamespaceDeleteRequest) (*wedo.NamespaceDeleteResponse, error) {
	panic("implement me")
}
