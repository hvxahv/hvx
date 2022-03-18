// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// BroadcastServiceClient is the client API for BroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastServiceClient interface {
	CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*CreateBroadcastResponse, error)
	GetBroadcasts(ctx context.Context, in *GetBroadcastsRequest, opts ...grpc.CallOption) (*GetBroadcastsResponse, error)
}

type broadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastServiceClient(cc grpc.ClientConnInterface) BroadcastServiceClient {
	return &broadcastServiceClient{cc}
}

func (c *broadcastServiceClient) CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*CreateBroadcastResponse, error) {
	out := new(CreateBroadcastResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.BroadcastService/CreateBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastServiceClient) GetBroadcasts(ctx context.Context, in *GetBroadcastsRequest, opts ...grpc.CallOption) (*GetBroadcastsResponse, error) {
	out := new(GetBroadcastsResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.BroadcastService/GetBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServiceServer is the server API for BroadcastService service.
// All implementations must embed UnimplementedBroadcastServiceServer
// for forward compatibility
type BroadcastServiceServer interface {
	CreateBroadcast(context.Context, *CreateBroadcastRequest) (*CreateBroadcastResponse, error)
	GetBroadcasts(context.Context, *GetBroadcastsRequest) (*GetBroadcastsResponse, error)
	mustEmbedUnimplementedBroadcastServiceServer()
}

// UnimplementedBroadcastServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBroadcastServiceServer struct {
}

func (UnimplementedBroadcastServiceServer) CreateBroadcast(context.Context, *CreateBroadcastRequest) (*CreateBroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBroadcast not implemented")
}
func (UnimplementedBroadcastServiceServer) GetBroadcasts(context.Context, *GetBroadcastsRequest) (*GetBroadcastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBroadcasts not implemented")
}
func (UnimplementedBroadcastServiceServer) mustEmbedUnimplementedBroadcastServiceServer() {}

// UnsafeBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServiceServer will
// result in compilation errors.
type UnsafeBroadcastServiceServer interface {
	mustEmbedUnimplementedBroadcastServiceServer()
}

func RegisterBroadcastServiceServer(s grpc.ServiceRegistrar, srv BroadcastServiceServer) {
	s.RegisterService(&BroadcastService_ServiceDesc, srv)
}

func _BroadcastService_CreateBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).CreateBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.BroadcastService/CreateBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).CreateBroadcast(ctx, req.(*CreateBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadcastService_GetBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBroadcastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).GetBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.BroadcastService/GetBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).GetBroadcasts(ctx, req.(*GetBroadcastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BroadcastService_ServiceDesc is the grpc.ServiceDesc for BroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.BroadcastService",
	HandlerType: (*BroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBroadcast",
			Handler:    _BroadcastService_CreateBroadcast_Handler,
		},
		{
			MethodName: "GetBroadcasts",
			Handler:    _BroadcastService_GetBroadcasts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/channel/v1alpha1/broadcast.proto",
}
