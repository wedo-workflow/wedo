// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wedo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WedoServiceClient is the client API for WedoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WedoServiceClient interface {
	// Get a deployment.
	DeploymentGet(ctx context.Context, in *DeploymentRequest, opts ...grpc.CallOption) (*DeploymentResponse, error)
	// Creates a deployment.
	DeploymentCreate(ctx context.Context, in *DeploymentCreateRequest, opts ...grpc.CallOption) (*DeploymentCreateResponse, error)
	// Updates a deployment.
	DeploymentList(ctx context.Context, in *DeploymentListRequest, opts ...grpc.CallOption) (*DeploymentListResponse, error)
	// Delete a deployment.
	DeploymentDelete(ctx context.Context, in *DeploymentDeleteRequest, opts ...grpc.CallOption) (*DeploymentDeleteResponse, error)
	// Create a User.
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	// Get a User.
	UserGet(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// Delete a User.
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	// List all Users.
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	UserListCount(ctx context.Context, in *UserListCountRequest, opts ...grpc.CallOption) (*UserListCountResponse, error)
	// Create a Namespace.
	NamespaceCreate(ctx context.Context, in *NamespaceCreateRequest, opts ...grpc.CallOption) (*NamespaceCreateResponse, error)
	// Get a Namespace.
	NamespaceGet(ctx context.Context, in *NamespaceRequest, opts ...grpc.CallOption) (*NamespaceResponse, error)
	// Delete a Namespace.
	NamespaceDelete(ctx context.Context, in *NamespaceDeleteRequest, opts ...grpc.CallOption) (*NamespaceDeleteResponse, error)
	// List all Namespaces.
	NamespaceList(ctx context.Context, in *NamespaceListRequest, opts ...grpc.CallOption) (*NamespaceListResponse, error)
	// Count all Namespaces.
	NamespaceListCount(ctx context.Context, in *NamespaceListCountRequest, opts ...grpc.CallOption) (*NamespaceListCountResponse, error)
	ProcessDefinitionStart(ctx context.Context, in *ProcessDefinitionCreateRequest, opts ...grpc.CallOption) (*ProcessDefinitionCreateResponse, error)
	ProcessDefinitionGet(ctx context.Context, in *ProcessDefinitionRequest, opts ...grpc.CallOption) (*ProcessDefinitionResponse, error)
	TaskCreate(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error)
	TaskGet(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	TaskDelete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
}

type wedoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWedoServiceClient(cc grpc.ClientConnInterface) WedoServiceClient {
	return &wedoServiceClient{cc}
}

func (c *wedoServiceClient) DeploymentGet(ctx context.Context, in *DeploymentRequest, opts ...grpc.CallOption) (*DeploymentResponse, error) {
	out := new(DeploymentResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) DeploymentCreate(ctx context.Context, in *DeploymentCreateRequest, opts ...grpc.CallOption) (*DeploymentCreateResponse, error) {
	out := new(DeploymentCreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) DeploymentList(ctx context.Context, in *DeploymentListRequest, opts ...grpc.CallOption) (*DeploymentListResponse, error) {
	out := new(DeploymentListResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) DeploymentDelete(ctx context.Context, in *DeploymentDeleteRequest, opts ...grpc.CallOption) (*DeploymentDeleteResponse, error) {
	out := new(DeploymentDeleteResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserGet(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) UserListCount(ctx context.Context, in *UserListCountRequest, opts ...grpc.CallOption) (*UserListCountResponse, error) {
	out := new(UserListCountResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserListCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) NamespaceCreate(ctx context.Context, in *NamespaceCreateRequest, opts ...grpc.CallOption) (*NamespaceCreateResponse, error) {
	out := new(NamespaceCreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) NamespaceGet(ctx context.Context, in *NamespaceRequest, opts ...grpc.CallOption) (*NamespaceResponse, error) {
	out := new(NamespaceResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) NamespaceDelete(ctx context.Context, in *NamespaceDeleteRequest, opts ...grpc.CallOption) (*NamespaceDeleteResponse, error) {
	out := new(NamespaceDeleteResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) NamespaceList(ctx context.Context, in *NamespaceListRequest, opts ...grpc.CallOption) (*NamespaceListResponse, error) {
	out := new(NamespaceListResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) NamespaceListCount(ctx context.Context, in *NamespaceListCountRequest, opts ...grpc.CallOption) (*NamespaceListCountResponse, error) {
	out := new(NamespaceListCountResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceListCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) ProcessDefinitionStart(ctx context.Context, in *ProcessDefinitionCreateRequest, opts ...grpc.CallOption) (*ProcessDefinitionCreateResponse, error) {
	out := new(ProcessDefinitionCreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/ProcessDefinitionStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) ProcessDefinitionGet(ctx context.Context, in *ProcessDefinitionRequest, opts ...grpc.CallOption) (*ProcessDefinitionResponse, error) {
	out := new(ProcessDefinitionResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/ProcessDefinitionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) TaskCreate(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskCreateResponse, error) {
	out := new(TaskCreateResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) TaskGet(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) TaskDelete(ctx context.Context, in *TaskDeleteRequest, opts ...grpc.CallOption) (*TaskDeleteResponse, error) {
	out := new(TaskDeleteResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wedoServiceClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WedoServiceServer is the server API for WedoService service.
// All implementations should embed UnimplementedWedoServiceServer
// for forward compatibility
type WedoServiceServer interface {
	// Get a deployment.
	DeploymentGet(context.Context, *DeploymentRequest) (*DeploymentResponse, error)
	// Creates a deployment.
	DeploymentCreate(context.Context, *DeploymentCreateRequest) (*DeploymentCreateResponse, error)
	// Updates a deployment.
	DeploymentList(context.Context, *DeploymentListRequest) (*DeploymentListResponse, error)
	// Delete a deployment.
	DeploymentDelete(context.Context, *DeploymentDeleteRequest) (*DeploymentDeleteResponse, error)
	// Create a User.
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	// Get a User.
	UserGet(context.Context, *UserRequest) (*UserResponse, error)
	// Delete a User.
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	// List all Users.
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	UserListCount(context.Context, *UserListCountRequest) (*UserListCountResponse, error)
	// Create a Namespace.
	NamespaceCreate(context.Context, *NamespaceCreateRequest) (*NamespaceCreateResponse, error)
	// Get a Namespace.
	NamespaceGet(context.Context, *NamespaceRequest) (*NamespaceResponse, error)
	// Delete a Namespace.
	NamespaceDelete(context.Context, *NamespaceDeleteRequest) (*NamespaceDeleteResponse, error)
	// List all Namespaces.
	NamespaceList(context.Context, *NamespaceListRequest) (*NamespaceListResponse, error)
	// Count all Namespaces.
	NamespaceListCount(context.Context, *NamespaceListCountRequest) (*NamespaceListCountResponse, error)
	ProcessDefinitionStart(context.Context, *ProcessDefinitionCreateRequest) (*ProcessDefinitionCreateResponse, error)
	ProcessDefinitionGet(context.Context, *ProcessDefinitionRequest) (*ProcessDefinitionResponse, error)
	TaskCreate(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error)
	TaskGet(context.Context, *TaskRequest) (*TaskResponse, error)
	TaskDelete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
}

// UnimplementedWedoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWedoServiceServer struct {
}

func (UnimplementedWedoServiceServer) DeploymentGet(context.Context, *DeploymentRequest) (*DeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentGet not implemented")
}
func (UnimplementedWedoServiceServer) DeploymentCreate(context.Context, *DeploymentCreateRequest) (*DeploymentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentCreate not implemented")
}
func (UnimplementedWedoServiceServer) DeploymentList(context.Context, *DeploymentListRequest) (*DeploymentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentList not implemented")
}
func (UnimplementedWedoServiceServer) DeploymentDelete(context.Context, *DeploymentDeleteRequest) (*DeploymentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentDelete not implemented")
}
func (UnimplementedWedoServiceServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedWedoServiceServer) UserGet(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (UnimplementedWedoServiceServer) UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedWedoServiceServer) UserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedWedoServiceServer) UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedWedoServiceServer) UserListCount(context.Context, *UserListCountRequest) (*UserListCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListCount not implemented")
}
func (UnimplementedWedoServiceServer) NamespaceCreate(context.Context, *NamespaceCreateRequest) (*NamespaceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceCreate not implemented")
}
func (UnimplementedWedoServiceServer) NamespaceGet(context.Context, *NamespaceRequest) (*NamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceGet not implemented")
}
func (UnimplementedWedoServiceServer) NamespaceDelete(context.Context, *NamespaceDeleteRequest) (*NamespaceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceDelete not implemented")
}
func (UnimplementedWedoServiceServer) NamespaceList(context.Context, *NamespaceListRequest) (*NamespaceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceList not implemented")
}
func (UnimplementedWedoServiceServer) NamespaceListCount(context.Context, *NamespaceListCountRequest) (*NamespaceListCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceListCount not implemented")
}
func (UnimplementedWedoServiceServer) ProcessDefinitionStart(context.Context, *ProcessDefinitionCreateRequest) (*ProcessDefinitionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDefinitionStart not implemented")
}
func (UnimplementedWedoServiceServer) ProcessDefinitionGet(context.Context, *ProcessDefinitionRequest) (*ProcessDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDefinitionGet not implemented")
}
func (UnimplementedWedoServiceServer) TaskCreate(context.Context, *TaskCreateRequest) (*TaskCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (UnimplementedWedoServiceServer) TaskGet(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskGet not implemented")
}
func (UnimplementedWedoServiceServer) TaskDelete(context.Context, *TaskDeleteRequest) (*TaskDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDelete not implemented")
}
func (UnimplementedWedoServiceServer) TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}

// UnsafeWedoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WedoServiceServer will
// result in compilation errors.
type UnsafeWedoServiceServer interface {
	mustEmbedUnimplementedWedoServiceServer()
}

func RegisterWedoServiceServer(s grpc.ServiceRegistrar, srv WedoServiceServer) {
	s.RegisterService(&WedoService_ServiceDesc, srv)
}

func _WedoService_DeploymentGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).DeploymentGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).DeploymentGet(ctx, req.(*DeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_DeploymentCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).DeploymentCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).DeploymentCreate(ctx, req.(*DeploymentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_DeploymentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).DeploymentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).DeploymentList(ctx, req.(*DeploymentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_DeploymentDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).DeploymentDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/DeploymentDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).DeploymentDelete(ctx, req.(*DeploymentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserGet(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_UserListCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).UserListCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/UserListCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).UserListCount(ctx, req.(*UserListCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_NamespaceCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).NamespaceCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).NamespaceCreate(ctx, req.(*NamespaceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_NamespaceGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).NamespaceGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).NamespaceGet(ctx, req.(*NamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_NamespaceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).NamespaceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).NamespaceDelete(ctx, req.(*NamespaceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_NamespaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).NamespaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).NamespaceList(ctx, req.(*NamespaceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_NamespaceListCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceListCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).NamespaceListCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/NamespaceListCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).NamespaceListCount(ctx, req.(*NamespaceListCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_ProcessDefinitionStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDefinitionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).ProcessDefinitionStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/ProcessDefinitionStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).ProcessDefinitionStart(ctx, req.(*ProcessDefinitionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_ProcessDefinitionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).ProcessDefinitionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/ProcessDefinitionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).ProcessDefinitionGet(ctx, req.(*ProcessDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).TaskCreate(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_TaskGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).TaskGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).TaskGet(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_TaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).TaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).TaskDelete(ctx, req.(*TaskDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WedoService_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WedoServiceServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.wedo_workflow.wedo.api.v1.WedoService/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WedoServiceServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WedoService_ServiceDesc is the grpc.ServiceDesc for WedoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WedoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.wedo_workflow.wedo.api.v1.WedoService",
	HandlerType: (*WedoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeploymentGet",
			Handler:    _WedoService_DeploymentGet_Handler,
		},
		{
			MethodName: "DeploymentCreate",
			Handler:    _WedoService_DeploymentCreate_Handler,
		},
		{
			MethodName: "DeploymentList",
			Handler:    _WedoService_DeploymentList_Handler,
		},
		{
			MethodName: "DeploymentDelete",
			Handler:    _WedoService_DeploymentDelete_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _WedoService_UserCreate_Handler,
		},
		{
			MethodName: "UserGet",
			Handler:    _WedoService_UserGet_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _WedoService_UserDelete_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _WedoService_UserList_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _WedoService_UserUpdate_Handler,
		},
		{
			MethodName: "UserListCount",
			Handler:    _WedoService_UserListCount_Handler,
		},
		{
			MethodName: "NamespaceCreate",
			Handler:    _WedoService_NamespaceCreate_Handler,
		},
		{
			MethodName: "NamespaceGet",
			Handler:    _WedoService_NamespaceGet_Handler,
		},
		{
			MethodName: "NamespaceDelete",
			Handler:    _WedoService_NamespaceDelete_Handler,
		},
		{
			MethodName: "NamespaceList",
			Handler:    _WedoService_NamespaceList_Handler,
		},
		{
			MethodName: "NamespaceListCount",
			Handler:    _WedoService_NamespaceListCount_Handler,
		},
		{
			MethodName: "ProcessDefinitionStart",
			Handler:    _WedoService_ProcessDefinitionStart_Handler,
		},
		{
			MethodName: "ProcessDefinitionGet",
			Handler:    _WedoService_ProcessDefinitionGet_Handler,
		},
		{
			MethodName: "TaskCreate",
			Handler:    _WedoService_TaskCreate_Handler,
		},
		{
			MethodName: "TaskGet",
			Handler:    _WedoService_TaskGet_Handler,
		},
		{
			MethodName: "TaskDelete",
			Handler:    _WedoService_TaskDelete_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _WedoService_TaskList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
