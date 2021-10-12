package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
)

func (APIServer) ProcessDefinitionStart(ctx context.Context, request *wedo.ProcessDefinitionCreateRequest) (*wedo.ProcessDefinitionCreateResponse, error) {
	panic("implement me")
}

func (APIServer) ProcessDefinitionGet(ctx context.Context, request *wedo.ProcessDefinitionRequest) (*wedo.ProcessDefinitionResponse, error) {
	panic("implement me")
}
