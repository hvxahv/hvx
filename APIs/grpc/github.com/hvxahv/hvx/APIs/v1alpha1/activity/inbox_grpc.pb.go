// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/activity/inbox.proto

package activity

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

// InboxClient is the client API for Inbox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InboxClient interface {
	Inbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxResponse, error)
	GetInbox(ctx context.Context, in *GetInboxRequest, opts ...grpc.CallOption) (*GetInboxResponse, error)
	GetInboxes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetInboxesResponse, error)
	DeleteInbox(ctx context.Context, in *DeleteInboxRequest, opts ...grpc.CallOption) (*DeleteInboxResponse, error)
	ViewedInbox(ctx context.Context, in *ViewedInboxRequest, opts ...grpc.CallOption) (*ViewedInboxResponse, error)
}

type inboxClient struct {
	cc grpc.ClientConnInterface
}

func NewInboxClient(cc grpc.ClientConnInterface) InboxClient {
	return &inboxClient{cc}
}

func (c *inboxClient) Inbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxResponse, error) {
	out := new(InboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Inbox/Inbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxClient) GetInbox(ctx context.Context, in *GetInboxRequest, opts ...grpc.CallOption) (*GetInboxResponse, error) {
	out := new(GetInboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Inbox/GetInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxClient) GetInboxes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetInboxesResponse, error) {
	out := new(GetInboxesResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Inbox/GetInboxes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxClient) DeleteInbox(ctx context.Context, in *DeleteInboxRequest, opts ...grpc.CallOption) (*DeleteInboxResponse, error) {
	out := new(DeleteInboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Inbox/DeleteInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inboxClient) ViewedInbox(ctx context.Context, in *ViewedInboxRequest, opts ...grpc.CallOption) (*ViewedInboxResponse, error) {
	out := new(ViewedInboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Inbox/ViewedInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InboxServer is the server API for Inbox service.
// All implementations should embed UnimplementedInboxServer
// for forward compatibility
type InboxServer interface {
	Inbox(context.Context, *InboxRequest) (*InboxResponse, error)
	GetInbox(context.Context, *GetInboxRequest) (*GetInboxResponse, error)
	GetInboxes(context.Context, *emptypb.Empty) (*GetInboxesResponse, error)
	DeleteInbox(context.Context, *DeleteInboxRequest) (*DeleteInboxResponse, error)
	ViewedInbox(context.Context, *ViewedInboxRequest) (*ViewedInboxResponse, error)
}

// UnimplementedInboxServer should be embedded to have forward compatible implementations.
type UnimplementedInboxServer struct {
}

func (UnimplementedInboxServer) Inbox(context.Context, *InboxRequest) (*InboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inbox not implemented")
}
func (UnimplementedInboxServer) GetInbox(context.Context, *GetInboxRequest) (*GetInboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInbox not implemented")
}
func (UnimplementedInboxServer) GetInboxes(context.Context, *emptypb.Empty) (*GetInboxesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboxes not implemented")
}
func (UnimplementedInboxServer) DeleteInbox(context.Context, *DeleteInboxRequest) (*DeleteInboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInbox not implemented")
}
func (UnimplementedInboxServer) ViewedInbox(context.Context, *ViewedInboxRequest) (*ViewedInboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewedInbox not implemented")
}

// UnsafeInboxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InboxServer will
// result in compilation errors.
type UnsafeInboxServer interface {
	mustEmbedUnimplementedInboxServer()
}

func RegisterInboxServer(s grpc.ServiceRegistrar, srv InboxServer) {
	s.RegisterService(&Inbox_ServiceDesc, srv)
}

func _Inbox_Inbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServer).Inbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Inbox/Inbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServer).Inbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inbox_GetInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServer).GetInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Inbox/GetInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServer).GetInbox(ctx, req.(*GetInboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inbox_GetInboxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServer).GetInboxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Inbox/GetInboxes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServer).GetInboxes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inbox_DeleteInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServer).DeleteInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Inbox/DeleteInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServer).DeleteInbox(ctx, req.(*DeleteInboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inbox_ViewedInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewedInboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InboxServer).ViewedInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Inbox/ViewedInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InboxServer).ViewedInbox(ctx, req.(*ViewedInboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inbox_ServiceDesc is the grpc.ServiceDesc for Inbox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inbox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.activity.proto.Inbox",
	HandlerType: (*InboxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inbox",
			Handler:    _Inbox_Inbox_Handler,
		},
		{
			MethodName: "GetInbox",
			Handler:    _Inbox_GetInbox_Handler,
		},
		{
			MethodName: "GetInboxes",
			Handler:    _Inbox_GetInboxes_Handler,
		},
		{
			MethodName: "DeleteInbox",
			Handler:    _Inbox_DeleteInbox_Handler,
		},
		{
			MethodName: "ViewedInbox",
			Handler:    _Inbox_ViewedInbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/activity/inbox.proto",
}
