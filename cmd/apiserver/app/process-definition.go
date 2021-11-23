package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) ProcessDefinitionStart(ctx context.Context, request *wedo.ProcessDefinitionStartRequest) (*wedo.ProcessDefinitionStartResponse, error) {
	if request.NamespaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "Namespace id is empty")
	}
	if request.ProcessDefinitionKey == "" {
		return nil, status.Error(codes.InvalidArgument, "Process definition id is empty")
	}

	return nil, nil
}

func (s *APIServer) ProcessDefinitionGet(ctx context.Context, request *wedo.ProcessDefinitionRequest) (*wedo.ProcessDefinitionResponse, error) {
	return nil, nil
}
