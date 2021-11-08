package app

import (
	"context"

	"github.com/google/uuid"
	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) UserCreate(ctx context.Context, request *wedo.UserCreateRequest) (*wedo.UserCreateResponse, error) {
	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if request.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is empty")
	}
	if request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is empty")
	}
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := &model.User{
		Id:       uuidV4.String(),
		Username: request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := s.Runtime.UserCreate(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.UserCreateResponse{
		Id:   user.Id,
		Name: request.Name,
	}, nil
}

func (s *APIServer) UserGet(ctx context.Context, request *wedo.UserRequest) (*wedo.UserResponse, error) {
	if request.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is empty")
	}
	user, err := s.Runtime.UserGet(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.UserResponse{
		Id:   request.UserId,
		Name: user.Username,
	}, nil
}

func (s *APIServer) UserDelete(ctx context.Context, request *wedo.UserDeleteRequest) (*wedo.UserDeleteResponse, error) {
	if request.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is empty")
	}
	if err := s.Runtime.UserDelete(ctx, request.UserId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &wedo.UserDeleteResponse{}, nil
}

func (APIServer) UserList(ctx context.Context, request *wedo.UserListRequest) (*wedo.UserListResponse, error) {
	panic("implement me")
}

func (APIServer) UserUpdate(ctx context.Context, request *wedo.UserUpdateRequest) (*wedo.UserUpdateResponse, error) {
	panic("implement me")
}

// UserListCount returns the count of the users.
func (s *APIServer) UserListCount(ctx context.Context, request *wedo.UserListCountRequest) (*wedo.UserListCountResponse, error) {
	panic("implement me")
}
