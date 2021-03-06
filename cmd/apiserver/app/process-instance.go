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
		ProcessInstanceId: request.ProcessInstanceId,
		State:             processInstance.Status,
	}, nil
}

func (s *APIServer) ProcessInstanceModify(ctx context.Context, request *wedo.ProcessModifyRequest) (*wedo.ProcessModifyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) ProcessInstanceActivateSuspend(ctx context.Context, request *wedo.ProcessActivateSuspendRequest) (*wedo.ProcessActivateSuspendResponse, error) {
	if request.ProcessInstanceId == "" {
		return nil, status.Error(codes.InvalidArgument, "ProcessInstanceId must be provided")
	}
	processInstance, err := s.Runtime.ProcessInstanceGet(ctx, &model.ProcessInstance{Id: request.ProcessInstanceId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var finalState = processInstance.Status
	if request.Suspend {
		finalState = wedo.ProcessInstanceState_SUSPENDED
	} else if finalState == wedo.ProcessInstanceState_SUSPENDED {
		finalState = wedo.ProcessInstanceState_STARTED
	}
	if err := s.Runtime.ProcessInstanceActivateSuspend(ctx, &model.ProcessInstance{
		Id:     request.ProcessInstanceId,
		Status: finalState,
	}); err != nil {
		return nil, err
	}
	return &wedo.ProcessActivateSuspendResponse{}, nil
}
