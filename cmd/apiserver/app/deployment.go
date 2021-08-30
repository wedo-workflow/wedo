package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
)

func (s *APIServer) DeploymentCreate(ctx context.Context, request *wedo.DeploymentCreateRequest) (*wedo.DeploymentCreateResponse, error) {
	return &wedo.DeploymentCreateResponse{}, nil
}
