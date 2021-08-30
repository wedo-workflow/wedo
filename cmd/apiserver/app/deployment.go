package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/element/wedo_model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) DeploymentCreate(ctx context.Context, request *wedo.DeploymentCreateRequest) (*wedo.DeploymentCreateResponse, error) {
	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if len(request.Content) == 0 {
		return nil, status.Error(codes.InvalidArgument, "content is empty")
	}
	if err := s.Runtime.Deploy(ctx, &wedo_model.Deploy{
		Name:    request.Name,
		Content: request.Content,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.DeploymentCreateResponse{}, nil
}
