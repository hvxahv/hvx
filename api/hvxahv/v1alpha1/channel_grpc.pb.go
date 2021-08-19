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
	// NewChannel Create Channel.
	NewChannel(ctx context.Context, in *NewChannelData, opts ...grpc.CallOption) (*NewChannelReply, error)
	// NewChannelAdmin Add Channel Admins.
	NewChannelAdmin(ctx context.Context, in *NewChanAdmData, opts ...grpc.CallOption) (*NewChanAdmReply, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) NewChannel(ctx context.Context, in *NewChannelData, opts ...grpc.CallOption) (*NewChannelReply, error) {
	out := new(NewChannelReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/NewChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) NewChannelAdmin(ctx context.Context, in *NewChanAdmData, opts ...grpc.CallOption) (*NewChanAdmReply, error) {
	out := new(NewChanAdmReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/NewChannelAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
// All implementations must embed UnimplementedChannelServer
// for forward compatibility
type ChannelServer interface {
	// NewChannel Create Channel.
	NewChannel(context.Context, *NewChannelData) (*NewChannelReply, error)
	// NewChannelAdmin Add Channel Admins.
	NewChannelAdmin(context.Context, *NewChanAdmData) (*NewChanAdmReply, error)
	mustEmbedUnimplementedChannelServer()
}

// UnimplementedChannelServer must be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (UnimplementedChannelServer) NewChannel(context.Context, *NewChannelData) (*NewChannelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChannel not implemented")
}
func (UnimplementedChannelServer) NewChannelAdmin(context.Context, *NewChanAdmData) (*NewChanAdmReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewChannelAdmin not implemented")
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

func _Channel_NewChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChannelData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).NewChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/NewChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).NewChannel(ctx, req.(*NewChannelData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_NewChannelAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChanAdmData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).NewChannelAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/NewChannelAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).NewChannelAdmin(ctx, req.(*NewChanAdmData))
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
			MethodName: "NewChannel",
			Handler:    _Channel_NewChannel_Handler,
		},
		{
			MethodName: "NewChannelAdmin",
			Handler:    _Channel_NewChannelAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hvxahv/v1alpha1/channel.proto",
}
