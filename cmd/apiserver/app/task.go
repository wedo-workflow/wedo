package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) TaskCreate(ctx context.Context, request *wedo.TaskCreateRequest) (*wedo.TaskCreateResponse, error) {
	id, err := s.Runtime.TaskCreate(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskCreateResponse{
		Id:   id,
		Name: request.Name,
	}, nil
}

func (s *APIServer) TaskGet(ctx context.Context, request *wedo.TaskRequest) (*wedo.TaskResponse, error) {
	if request.TaskId == "" {
		return nil, status.Error(codes.InvalidArgument, "task id is empty")
	}
	task, err := s.Runtime.TaskGet(ctx, request.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return task, nil
}

func (s *APIServer) TaskDelete(ctx context.Context, request *wedo.TaskDeleteRequest) (*wedo.TaskDeleteResponse, error) {
	if request.TaskId == "" {
		return nil, status.Error(codes.InvalidArgument, "task id is empty")
	}
	err := s.Runtime.TaskDelete(ctx, request.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskDeleteResponse{}, nil
}

// TaskList lists all tasks.
func (s *APIServer) TaskList(ctx context.Context, request *wedo.TaskListRequest) (*wedo.TaskListResponse, error) {
	if request.ProcessInstanceId == "" {
		return nil, status.Error(codes.InvalidArgument, "process instance id is empty")
	}
	tasks, err := s.Runtime.TaskList(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.TaskListResponse{
		Tasks: tasks,
	}, nil
}

func (s *APIServer) TaskClaim(ctx context.Context, request *wedo.TaskClaimRequest) (*wedo.TaskClaimResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskUnclaim(ctx context.Context, request *wedo.TaskUnclaimRequest) (*wedo.TaskUnclaimResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskComplete(ctx context.Context, request *wedo.TaskCompleteRequest) (*wedo.TaskCompleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskGetFormKey(ctx context.Context, request *wedo.TaskGetFormKeyRequest) (*wedo.TaskGetFormKeyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskDelegate(ctx context.Context, request *wedo.TaskDelegateRequest) (*wedo.TaskDelegateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskAssign(ctx context.Context, request *wedo.TaskAssignRequest) (*wedo.TaskAssignResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskSubmitForm(ctx context.Context, request *wedo.TaskSubmitFormRequest) (*wedo.TaskSubmitFormResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskResolve(ctx context.Context, request *wedo.TaskResolveRequest) (*wedo.TaskResolveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *APIServer) TaskDeployedForm(ctx context.Context, request *wedo.TaskDeployedFormRequest) (*wedo.TaskDeployedFormResponse, error) {
	//TODO implement me
	panic("implement me")
}
