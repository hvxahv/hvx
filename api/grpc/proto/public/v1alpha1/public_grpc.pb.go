// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PublicClient is the client API for Public service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicClient interface {
	// Get the instance details of the current instance.
	GetPublicInstance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicInstanceResponse, error)
	// Get the total number of users of the current instance.
	GetPublicAccountCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicAccountCountResponse, error)
	CreateAccounts(ctx context.Context, in *CreateAccountsRequest, opts ...grpc.CallOption) (*CreateAccountsResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	GetWebfinger(ctx context.Context, in *GetWebfingerRequest, opts ...grpc.CallOption) (*GetWebfingerResponse, error)
	// Get the actors in the activityPub protocol.
	// https://www.w3.org/TR/activitypub/#actor-objects
	GetActor(ctx context.Context, in *GetActorRequest, opts ...grpc.CallOption) (*GetActorResponse, error)
}

type publicClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicClient(cc grpc.ClientConnInterface) PublicClient {
	return &publicClient{cc}
}

func (c *publicClient) GetPublicInstance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicInstanceResponse, error) {
	out := new(GetPublicInstanceResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/GetPublicInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetPublicAccountCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicAccountCountResponse, error) {
	out := new(GetPublicAccountCountResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/GetPublicAccountCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) CreateAccounts(ctx context.Context, in *CreateAccountsRequest, opts ...grpc.CallOption) (*CreateAccountsResponse, error) {
	out := new(CreateAccountsResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/CreateAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetWebfinger(ctx context.Context, in *GetWebfingerRequest, opts ...grpc.CallOption) (*GetWebfingerResponse, error) {
	out := new(GetWebfingerResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/GetWebfinger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) GetActor(ctx context.Context, in *GetActorRequest, opts ...grpc.CallOption) (*GetActorResponse, error) {
	out := new(GetActorResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Public/GetActor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServer is the server API for Public service.
// All implementations should embed UnimplementedPublicServer
// for forward compatibility
type PublicServer interface {
	// Get the instance details of the current instance.
	GetPublicInstance(context.Context, *emptypb.Empty) (*GetPublicInstanceResponse, error)
	// Get the total number of users of the current instance.
	GetPublicAccountCount(context.Context, *emptypb.Empty) (*GetPublicAccountCountResponse, error)
	CreateAccounts(context.Context, *CreateAccountsRequest) (*CreateAccountsResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	GetWebfinger(context.Context, *GetWebfingerRequest) (*GetWebfingerResponse, error)
	// Get the actors in the activityPub protocol.
	// https://www.w3.org/TR/activitypub/#actor-objects
	GetActor(context.Context, *GetActorRequest) (*GetActorResponse, error)
}

// UnimplementedPublicServer should be embedded to have forward compatible implementations.
type UnimplementedPublicServer struct {
}

func (UnimplementedPublicServer) GetPublicInstance(context.Context, *emptypb.Empty) (*GetPublicInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicInstance not implemented")
}
func (UnimplementedPublicServer) GetPublicAccountCount(context.Context, *emptypb.Empty) (*GetPublicAccountCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicAccountCount not implemented")
}
func (UnimplementedPublicServer) CreateAccounts(context.Context, *CreateAccountsRequest) (*CreateAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccounts not implemented")
}
func (UnimplementedPublicServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedPublicServer) GetWebfinger(context.Context, *GetWebfingerRequest) (*GetWebfingerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebfinger not implemented")
}
func (UnimplementedPublicServer) GetActor(context.Context, *GetActorRequest) (*GetActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActor not implemented")
}

// UnsafePublicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicServer will
// result in compilation errors.
type UnsafePublicServer interface {
	mustEmbedUnimplementedPublicServer()
}

func RegisterPublicServer(s grpc.ServiceRegistrar, srv PublicServer) {
	s.RegisterService(&Public_ServiceDesc, srv)
}

func _Public_GetPublicInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetPublicInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/GetPublicInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetPublicInstance(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetPublicAccountCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetPublicAccountCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/GetPublicAccountCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetPublicAccountCount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_CreateAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).CreateAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/CreateAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).CreateAccounts(ctx, req.(*CreateAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetWebfinger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWebfingerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetWebfinger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/GetWebfinger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetWebfinger(ctx, req.(*GetWebfingerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_GetActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).GetActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Public/GetActor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).GetActor(ctx, req.(*GetActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Public_ServiceDesc is the grpc.ServiceDesc for Public service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Public_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Public",
	HandlerType: (*PublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicInstance",
			Handler:    _Public_GetPublicInstance_Handler,
		},
		{
			MethodName: "GetPublicAccountCount",
			Handler:    _Public_GetPublicAccountCount_Handler,
		},
		{
			MethodName: "CreateAccounts",
			Handler:    _Public_CreateAccounts_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Public_Authenticate_Handler,
		},
		{
			MethodName: "GetWebfinger",
			Handler:    _Public_GetWebfinger_Handler,
		},
		{
			MethodName: "GetActor",
			Handler:    _Public_GetActor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/public/v1alpha1/public.proto",
}
