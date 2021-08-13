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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	// NewAccount Create an account through the NewAccountData data structure.
	NewAccount(ctx context.Context, in *NewAccountData, opts ...grpc.CallOption) (*AccountsReply, error)
	// UpdateAccount Update the account through AccountData data.
	UpdateAccount(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountsReply, error)
	// DeleteAccount Delete the specified account by user name.
	DeleteAccount(ctx context.Context, in *AccountByName, opts ...grpc.CallOption) (*AccountsReply, error)
	// FindAccount Query specified user data by username.
	FindAccount(ctx context.Context, in *AccountByName, opts ...grpc.CallOption) (*AccountData, error)
	LoginAccount(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*LoginReply, error)
	//  New follower.
	NewFollow(ctx context.Context, in *FollowersData, opts ...grpc.CallOption) (*AccountsReply, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) NewAccount(ctx context.Context, in *NewAccountData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/NewAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) UpdateAccount(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) DeleteAccount(ctx context.Context, in *AccountByName, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) FindAccount(ctx context.Context, in *AccountByName, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/FindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) LoginAccount(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/LoginAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) NewFollow(ctx context.Context, in *FollowersData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/NewFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// NewAccount Create an account through the NewAccountData data structure.
	NewAccount(context.Context, *NewAccountData) (*AccountsReply, error)
	// UpdateAccount Update the account through AccountData data.
	UpdateAccount(context.Context, *AccountData) (*AccountsReply, error)
	// DeleteAccount Delete the specified account by user name.
	DeleteAccount(context.Context, *AccountByName) (*AccountsReply, error)
	// FindAccount Query specified user data by username.
	FindAccount(context.Context, *AccountByName) (*AccountData, error)
	LoginAccount(context.Context, *LoginData) (*LoginReply, error)
	//  New follower.
	NewFollow(context.Context, *FollowersData) (*AccountsReply, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) NewAccount(context.Context, *NewAccountData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAccount not implemented")
}
func (UnimplementedAccountsServer) UpdateAccount(context.Context, *AccountData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountsServer) DeleteAccount(context.Context, *AccountByName) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountsServer) FindAccount(context.Context, *AccountByName) (*AccountData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccount not implemented")
}
func (UnimplementedAccountsServer) LoginAccount(context.Context, *LoginData) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAccount not implemented")
}
func (UnimplementedAccountsServer) NewFollow(context.Context, *FollowersData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewFollow not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s grpc.ServiceRegistrar, srv AccountsServer) {
	s.RegisterService(&Accounts_ServiceDesc, srv)
}

func _Accounts_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/NewAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).NewAccount(ctx, req.(*NewAccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).UpdateAccount(ctx, req.(*AccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).DeleteAccount(ctx, req.(*AccountByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_FindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).FindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/FindAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).FindAccount(ctx, req.(*AccountByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_LoginAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).LoginAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/LoginAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).LoginAccount(ctx, req.(*LoginData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_NewFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowersData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).NewFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/NewFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).NewFollow(ctx, req.(*FollowersData))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAccount",
			Handler:    _Accounts_NewAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Accounts_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Accounts_DeleteAccount_Handler,
		},
		{
			MethodName: "FindAccount",
			Handler:    _Accounts_FindAccount_Handler,
		},
		{
			MethodName: "LoginAccount",
			Handler:    _Accounts_LoginAccount_Handler,
		},
		{
			MethodName: "NewFollow",
			Handler:    _Accounts_NewFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hvxahv/v1alpha1/accounts.proto",
}
