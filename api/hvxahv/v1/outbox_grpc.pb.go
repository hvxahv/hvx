// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// OutboxClient is the client API for Outbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboxClient interface {
	Accept(ctx context.Context, in *AcceptData, opts ...grpc.CallOption) (*ReplyCode, error)
	Follow(ctx context.Context, in *FollowData, opts ...grpc.CallOption) (*ReplyCode, error)
}

type outboxClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboxClient(cc grpc.ClientConnInterface) OutboxClient {
	return &outboxClient{cc}
}

func (c *outboxClient) Accept(ctx context.Context, in *AcceptData, opts ...grpc.CallOption) (*ReplyCode, error) {
	out := new(ReplyCode)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Outbox/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxClient) Follow(ctx context.Context, in *FollowData, opts ...grpc.CallOption) (*ReplyCode, error) {
	out := new(ReplyCode)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Outbox/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutboxServer is the server API for Outbox service.
// All implementations must embed UnimplementedOutboxServer
// for forward compatibility
type OutboxServer interface {
	Accept(context.Context, *AcceptData) (*ReplyCode, error)
	Follow(context.Context, *FollowData) (*ReplyCode, error)
	mustEmbedUnimplementedOutboxServer()
}

// UnimplementedOutboxServer must be embedded to have forward compatible implementations.
type UnimplementedOutboxServer struct {
}

func (UnimplementedOutboxServer) Accept(context.Context, *AcceptData) (*ReplyCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (UnimplementedOutboxServer) Follow(context.Context, *FollowData) (*ReplyCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedOutboxServer) mustEmbedUnimplementedOutboxServer() {}

// UnsafeOutboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboxServer will
// result in compilation errors.
type UnsafeOutboxServer interface {
	mustEmbedUnimplementedOutboxServer()
}

func RegisterOutboxServer(s grpc.ServiceRegistrar, srv OutboxServer) {
	s.RegisterService(&Outbox_ServiceDesc, srv)
}

func _Outbox_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Outbox/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServer).Accept(ctx, req.(*AcceptData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Outbox_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Outbox/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServer).Follow(ctx, req.(*FollowData))
	}
	return interceptor(ctx, in, info, handler)
}

// Outbox_ServiceDesc is the grpc.ServiceDesc for Outbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Outbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1.proto.Outbox",
	HandlerType: (*OutboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Accept",
			Handler:    _Outbox_Accept_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _Outbox_Follow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hvxahv/v1/outbox.proto",
}
