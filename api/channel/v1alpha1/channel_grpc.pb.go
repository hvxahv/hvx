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

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelClient interface {
	// New Create Channel.
	New(ctx context.Context, in *NewChannelData, opts ...grpc.CallOption) (*NewChannelReply, error)
	// NewAdmin Add Channel Admins.
	NewAdmin(ctx context.Context, in *NewAdminData, opts ...grpc.CallOption) (*ChannelReply, error)
	// NewSubscribers New Subscriber.
	NewSubscriber(ctx context.Context, in *NewSubscriberData, opts ...grpc.CallOption) (*ChannelReply, error)
	// GetSubscribers Get channels subscribers.
	GetSubscribers(ctx context.Context, in *NewChannelByID, opts ...grpc.CallOption) (*GetSubscribersReply, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) New(ctx context.Context, in *NewChannelData, opts ...grpc.CallOption) (*NewChannelReply, error) {
	out := new(NewChannelReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) NewAdmin(ctx context.Context, in *NewAdminData, opts ...grpc.CallOption) (*ChannelReply, error) {
	out := new(ChannelReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/NewAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) NewSubscriber(ctx context.Context, in *NewSubscriberData, opts ...grpc.CallOption) (*ChannelReply, error) {
	out := new(ChannelReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/NewSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) GetSubscribers(ctx context.Context, in *NewChannelByID, opts ...grpc.CallOption) (*GetSubscribersReply, error) {
	out := new(GetSubscribersReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/GetSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
// All implementations must embed UnimplementedChannelServer
// for forward compatibility
type ChannelServer interface {
	// New Create Channel.
	New(context.Context, *NewChannelData) (*NewChannelReply, error)
	// NewAdmin Add Channel Admins.
	NewAdmin(context.Context, *NewAdminData) (*ChannelReply, error)
	// NewSubscribers New Subscriber.
	NewSubscriber(context.Context, *NewSubscriberData) (*ChannelReply, error)
	// GetSubscribers Get channels subscribers.
	GetSubscribers(context.Context, *NewChannelByID) (*GetSubscribersReply, error)
	mustEmbedUnimplementedChannelServer()
}

// UnimplementedChannelServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (UnimplementedChannelServer) New(context.Context, *NewChannelData) (*NewChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedChannelServer) NewAdmin(context.Context, *NewAdminData) (*ChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAdmin not implemented")
}
func (UnimplementedChannelServer) NewSubscriber(context.Context, *NewSubscriberData) (*ChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSubscriber not implemented")
}
func (UnimplementedChannelServer) GetSubscribers(context.Context, *NewChannelByID) (*GetSubscribersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribers not implemented")
}
func (UnimplementedChannelServer) mustEmbedUnimplementedChannelServer() {}

// UnsafeChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServer will
// result in compilation errors.
type UnsafeChannelServer interface {
	mustEmbedUnimplementedChannelServer()
}

func RegisterChannelServer(s grpc.ServiceRegistrar, srv ChannelServer) {
	s.RegisterService(&Channel_ServiceDesc, srv)
}

func _Channel_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChannelData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).New(ctx, req.(*NewChannelData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_NewAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAdminData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).NewAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/NewAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).NewAdmin(ctx, req.(*NewAdminData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_NewSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSubscriberData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).NewSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/NewSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).NewSubscriber(ctx, req.(*NewSubscriberData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_GetSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChannelByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).GetSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/GetSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).GetSubscribers(ctx, req.(*NewChannelByID))
	}
	return interceptor(ctx, in, info, handler)
}

// Channel_ServiceDesc is the grpc.ServiceDesc for Channel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Channel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Channel_New_Handler,
		},
		{
			MethodName: "NewAdmin",
			Handler:    _Channel_NewAdmin_Handler,
		},
		{
			MethodName: "NewSubscriber",
			Handler:    _Channel_NewSubscriber_Handler,
		},
		{
			MethodName: "GetSubscribers",
			Handler:    _Channel_GetSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/channels/v1alpha1/channels.proto",
}
