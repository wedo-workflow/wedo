package app

import (
	"context"
	"time"

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
	if request.NamespaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "namespace is empty")
	}
	now := time.Now()
	deploy, err := s.Runtime.Deploy(ctx, &model.Deployment{
		NamespaceID: request.NamespaceId,
		Name:        request.Name,
		Content:     request.Content,
		CreateTime:  now,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentCreateResponse{
		Id:                  deploy.DID,
		Name:                request.Name,
		Content:             request.Content,
		Timestamp:           timestamppb.New(now),
		NamespaceId:         request.NamespaceId,
		ProcessDefinitionId: deploy.ProcessDefinitionID,
	}
	return ret, nil
}

func (s *APIServer) DeploymentGet(ctx context.Context, request *wedo.DeploymentRequest) (*wedo.DeploymentResponse, error) {
	if request.DeploymentId == "" {
		return nil, status.Error(codes.InvalidArgument, "deployment id is empty")
	}
	deploy, err := s.Runtime.Store().Deployment(ctx, request.DeploymentId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentResponse{
		Id:                  request.DeploymentId,
		Name:                deploy.Name,
		Content:             deploy.Content,
		Timestamp:           timestamppb.New(deploy.CreateTime),
		ProcessDefinitionId: deploy.ProcessDefinitionID,
	}
	return ret, nil
}

// DeploymentList returns a list of deployments.
func (s *APIServer) DeploymentList(ctx context.Context, request *wedo.DeploymentListRequest) (*wedo.DeploymentListResponse, error) {
	deploys, err := s.Runtime.DeployList(ctx, &model.DeploymentListOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ret := &wedo.DeploymentListResponse{
		Deployments: make([]*wedo.DeploymentResponse, 0),
	}
	for _, deploy := range deploys {
		ret.Deployments = append(ret.Deployments, &wedo.DeploymentResponse{
			Id:                  deploy.DID,
			Name:                deploy.Name,
			Content:             deploy.Content,
			Timestamp:           timestamppb.New(deploy.CreateTime),
			ProcessDefinitionId: deploy.ProcessDefinitionID,
		})
	}
	return ret, nil
}

// DeploymentDelete deletes a deployment.
func (s *APIServer) DeploymentDelete(ctx context.Context, request *wedo.DeploymentDeleteRequest) (*wedo.DeploymentDeleteResponse, error) {
	if request.DeploymentId == "" {
		return nil, status.Error(codes.InvalidArgument, "deployment id is empty")
	}
	err := s.Runtime.DeployDelete(ctx, request.DeploymentId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.DeploymentDeleteResponse{}, nil
}
