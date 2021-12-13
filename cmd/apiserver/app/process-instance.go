package app

import (
	"context"
	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) ProcessInstanceGet(ctx context.Context, request *wedo.ProcessInstanceRequest) (*wedo.ProcessInstanceResponse, error) {
	if request.ProcessInstanceId == "" {
		return nil, status.Error(codes.InvalidArgument, "ProcessInstanceId must be provided")
	}
	processInstance, err := s.Runtime.ProcessInstanceGet(ctx, &model.ProcessInstance{Id: request.ProcessInstanceId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.ProcessInstanceResponse{
		ProcessInstanceId:    request.ProcessInstanceId,
		ProcessDefinitionKey: "",
		NamespaceId:          "",
		State:                processInstance.Status,
	}, nil
}

func (s *APIServer) ProcessModify(ctx context.Context, request *wedo.ProcessModifyRequest) (*wedo.ProcessModifyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) ProcessActivateSuspend(ctx context.Context, request *wedo.ProcessActivateSuspendRequest) (*wedo.ProcessActivateSuspendResponse, error) {
	//TODO implement me
	panic("implement me")
}
