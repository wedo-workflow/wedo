package app

import (
	"context"

	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *APIServer) NamespaceCreate(ctx context.Context, request *wedo.NamespaceCreateRequest) (*wedo.NamespaceCreateResponse, error) {
	if request.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	namespce, err := s.Runtime.NamespaceCreate(ctx, &model.Namespace{
		Name: request.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := &wedo.NamespaceCreateResponse{
		Id:   namespce.ID,
		Name: request.Name,
	}
	return res, err
}

func (s *APIServer) NamespaceGet(ctx context.Context, request *wedo.NamespaceRequest) (*wedo.NamespaceResponse, error) {
	if request.NamespaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	namespace, err := s.Runtime.NamespaceGet(ctx, request.NamespaceId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "namespace not found")
	}
	res := &wedo.NamespaceResponse{
		Id:   namespace.ID,
		Name: namespace.Name,
	}
	return res, err
}

// NamespaceList returns a list of namespaces
func (s *APIServer) NamespaceList(ctx context.Context, request *wedo.NamespaceListRequest) (*wedo.NamespaceListResponse, error) {
	namespaces, err := s.Runtime.NamespaceList(ctx, &model.NamespaceQueryOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := &wedo.NamespaceListResponse{
		Namespaces: make([]*wedo.NamespaceResponse, 0, len(namespaces)),
	}
	for _, namespace := range namespaces {
		res.Namespaces = append(res.Namespaces, &wedo.NamespaceResponse{
			Id:   namespace.ID,
			Name: namespace.Name,
		})
	}
	return res, err
}

// NamespaceListCount returns a count of namespaces
func (s *APIServer) NamespaceListCount(ctx context.Context, request *wedo.NamespaceListCountRequest) (*wedo.NamespaceListCountResponse, error) {
	count, err := s.Runtime.NamespaceListCount(ctx, &model.NamespaceQueryOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := &wedo.NamespaceListCountResponse{
		Count: count,
	}
	return res, err
}

func (s *APIServer) NamespaceDelete(ctx context.Context, request *wedo.NamespaceDeleteRequest) (*wedo.NamespaceDeleteResponse, error) {
	if request.NamespaceId == "" {
		return nil, status.Error(codes.InvalidArgument, "id is empty")
	}
	err := s.Runtime.NamespaceDelete(ctx, request.NamespaceId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "namespace not found")
	}
	res := &wedo.NamespaceDeleteResponse{}
	return res, err
}
