// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/channel/broadcast.proto

package channel

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

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastClient interface {
	CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*CreateBroadcastResponse, error)
	GetBroadcasts(ctx context.Context, in *GetBroadcastsRequest, opts ...grpc.CallOption) (*GetBroadcastsResponse, error)
	DeleteBroadcast(ctx context.Context, in *DeleteBroadcastRequest, opts ...grpc.CallOption) (*DeleteBroadcastResponse, error)
}

type broadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastClient(cc grpc.ClientConnInterface) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) CreateBroadcast(ctx context.Context, in *CreateBroadcastRequest, opts ...grpc.CallOption) (*CreateBroadcastResponse, error) {
	out := new(CreateBroadcastResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.channel.proto.Broadcast/CreateBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) GetBroadcasts(ctx context.Context, in *GetBroadcastsRequest, opts ...grpc.CallOption) (*GetBroadcastsResponse, error) {
	out := new(GetBroadcastsResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.channel.proto.Broadcast/GetBroadcasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) DeleteBroadcast(ctx context.Context, in *DeleteBroadcastRequest, opts ...grpc.CallOption) (*DeleteBroadcastResponse, error) {
	out := new(DeleteBroadcastResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.channel.proto.Broadcast/DeleteBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
// All implementations should embed UnimplementedBroadcastServer
// for forward compatibility
type BroadcastServer interface {
	CreateBroadcast(context.Context, *CreateBroadcastRequest) (*CreateBroadcastResponse, error)
	GetBroadcasts(context.Context, *GetBroadcastsRequest) (*GetBroadcastsResponse, error)
	DeleteBroadcast(context.Context, *DeleteBroadcastRequest) (*DeleteBroadcastResponse, error)
}

// UnimplementedBroadcastServer should be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (UnimplementedBroadcastServer) CreateBroadcast(context.Context, *CreateBroadcastRequest) (*CreateBroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBroadcast not implemented")
}
func (UnimplementedBroadcastServer) GetBroadcasts(context.Context, *GetBroadcastsRequest) (*GetBroadcastsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBroadcasts not implemented")
}
func (UnimplementedBroadcastServer) DeleteBroadcast(context.Context, *DeleteBroadcastRequest) (*DeleteBroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBroadcast not implemented")
}

// UnsafeBroadcastServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServer will
// result in compilation errors.
type UnsafeBroadcastServer interface {
	mustEmbedUnimplementedBroadcastServer()
}

func RegisterBroadcastServer(s grpc.ServiceRegistrar, srv BroadcastServer) {
	s.RegisterService(&Broadcast_ServiceDesc, srv)
}

func _Broadcast_CreateBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).CreateBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.channel.proto.Broadcast/CreateBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).CreateBroadcast(ctx, req.(*CreateBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_GetBroadcasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBroadcastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).GetBroadcasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.channel.proto.Broadcast/GetBroadcasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).GetBroadcasts(ctx, req.(*GetBroadcastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_DeleteBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).DeleteBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.channel.proto.Broadcast/DeleteBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).DeleteBroadcast(ctx, req.(*DeleteBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Broadcast_ServiceDesc is the grpc.ServiceDesc for Broadcast service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broadcast_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.channel.proto.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBroadcast",
			Handler:    _Broadcast_CreateBroadcast_Handler,
		},
		{
			MethodName: "GetBroadcasts",
			Handler:    _Broadcast_GetBroadcasts_Handler,
		},
		{
			MethodName: "DeleteBroadcast",
			Handler:    _Broadcast_DeleteBroadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/channel/broadcast.proto",
}
