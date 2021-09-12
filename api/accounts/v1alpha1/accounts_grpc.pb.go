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
	// New Create an account through the NewAccountData data structure.
	New(ctx context.Context, in *NewAccountData, opts ...grpc.CallOption) (*AccountsReply, error)
	// Update Update the account through AccountData data.
	Update(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountsReply, error)
	// Delete Delete the specified account by user name.
	Delete(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AccountsReply, error)
	// Query account specified user data by username.
	QueryByName(ctx context.Context, in *NewAccountByName, opts ...grpc.CallOption) (*AccountData, error)
	// Query account specified user data by id.
	QueryByID(ctx context.Context, in *NewAccountByID, opts ...grpc.CallOption) (*AccountData, error)
	// Login ...
	Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AuthReply, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) New(ctx context.Context, in *NewAccountData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Update(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Delete(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AccountsReply, error) {
	out := new(AccountsReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) QueryByName(ctx context.Context, in *NewAccountByName, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/QueryByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) QueryByID(ctx context.Context, in *NewAccountByID, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/QueryByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Login(ctx context.Context, in *AuthData, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// New Create an account through the NewAccountData data structure.
	New(context.Context, *NewAccountData) (*AccountsReply, error)
	// Update Update the account through AccountData data.
	Update(context.Context, *AccountData) (*AccountsReply, error)
	// Delete Delete the specified account by user name.
	Delete(context.Context, *AuthData) (*AccountsReply, error)
	// Query account specified user data by username.
	QueryByName(context.Context, *NewAccountByName) (*AccountData, error)
	// Query account specified user data by id.
	QueryByID(context.Context, *NewAccountByID) (*AccountData, error)
	// Login ...
	Login(context.Context, *AuthData) (*AuthReply, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) New(context.Context, *NewAccountData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedAccountsServer) Update(context.Context, *AccountData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccountsServer) Delete(context.Context, *AuthData) (*AccountsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountsServer) QueryByName(context.Context, *NewAccountByName) (*AccountData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryByName not implemented")
}
func (UnimplementedAccountsServer) QueryByID(context.Context, *NewAccountByID) (*AccountData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryByID not implemented")
}
func (UnimplementedAccountsServer) Login(context.Context, *AuthData) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
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

func _Accounts_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).New(ctx, req.(*NewAccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Update(ctx, req.(*AccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Delete(ctx, req.(*AuthData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_QueryByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).QueryByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/QueryByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).QueryByName(ctx, req.(*NewAccountByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_QueryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).QueryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/QueryByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).QueryByID(ctx, req.(*NewAccountByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Login(ctx, req.(*AuthData))
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
			MethodName: "New",
			Handler:    _Accounts_New_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Accounts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Accounts_Delete_Handler,
		},
		{
			MethodName: "QueryByName",
			Handler:    _Accounts_QueryByName_Handler,
		},
		{
			MethodName: "QueryByID",
			Handler:    _Accounts_QueryByID_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Accounts_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/accounts/v1alpha1/accounts.proto",
}
