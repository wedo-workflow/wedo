package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
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
	if request.NamespaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "Namespace id is empty")
	}
	if request.ProcessDefinitionKey == "" {
		return nil, status.Error(codes.InvalidArgument, "Process definition id is empty")
	}
	rt, err := s.Runtime.ProcessDefinitionGet(ctx, &model.ProcessDefinition{
		ProcessDefinitionKey: request.ProcessDefinitionKey,
		NamespaceId:          request.NamespaceId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.ProcessDefinitionResponse{
		ProcessDefinitionKey: rt.BusinessName,
		NamespaceId:          rt.NamespaceID,
	}, nil
}
