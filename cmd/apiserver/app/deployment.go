package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *APIServer) DeploymentCreate(ctx context.Context, request *wedo.DeploymentCreateRequest) (*wedo.DeploymentCreateResponse, error) {
	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if len(request.Content) == 0 {
		return nil, status.Error(codes.InvalidArgument, "content is empty")
	}
	id, err := s.Runtime.Deploy(ctx, &model.Deploy{
		Name:    request.Name,
		Content: request.Content,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentCreateResponse{
		Id:      id,
		Name:    request.Name,
		Content: request.Content,
	}
	return ret, nil
}

func (s *APIServer) DeploymentGet(ctx context.Context, request *wedo.DeploymentRequest) (*wedo.DeploymentResponse, error) {
	if request.DeploymentId == "" {
		return nil, status.Error(codes.InvalidArgument, "deployment id is empty")
	}
	deploy, err := s.Runtime.Store().Deploy(ctx, request.DeploymentId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentResponse{
		Id:        request.DeploymentId,
		Name:      deploy.Name,
		Content:   deploy.Content,
		Timestamp: timestamppb.New(deploy.CreateTime),
	}
	return ret, nil
}

// DeploymentList returns a list of deployments.
func (s *APIServer) DeploymentList(ctx context.Context, request *wedo.DeploymentListRequest) (*wedo.DeploymentListResponse, error) {
	deploys, err := s.Runtime.Store().DeployList(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentListResponse{
		Deployments: make([]*wedo.DeploymentResponse, 0, len(deploys)),
	}
	for _, deploy := range deploys {
		ret.Deployments = append(ret.Deployments, &wedo.DeploymentResponse{
			Id:        deploy.Id,
			Name:      deploy.Name,
			Content:   deploy.Content,
			Timestamp: timestamppb.New(deploy.CreateTime),
		})
	}
	return ret, nil
}
