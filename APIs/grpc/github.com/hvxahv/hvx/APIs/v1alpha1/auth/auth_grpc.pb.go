// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/auth/auth.proto

package auth

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// Verify authentication Login with a valid user.
	// After successful authentication. The client must be registered to the device table for functions such as TOKEN checksum, managing devices, etc.
	// So you must submit a UA identifier in addition to the username and password when logging in.
	// A valid Token is returned and must be carried in subsequent API access operations.
	// https://datatracker.ietf.org/doc/html/rfc9068
	Authorization(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error)
	// SetPublicKey unlike activitypub, his private key is inaccessible to the server.
	// The public key is used for hvxahv privacy-related asymmetric encryption key.
	SetPublicKey(ctx context.Context, in *SetPublicKeyRequest, opts ...grpc.CallOption) (*SetPublicKeyResponse, error)
	// GetPublicKey Use the account id to get the account public key.
	// Not activitypub public key.
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
	// GetPrivateKey First, initiate a request to obtain the private key from the logged-in client.
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error)
	// GetDH Get the dh parameter.
	GetDH(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDHResponse, error)
	// SendPrivateKey Send the encrypted private key.
	SendPrivateKey(ctx context.Context, in *SendPrivateKeyRequest, opts ...grpc.CallOption) (*SendPrivateKeyResponse, error)
	// WaitPrivateKey Wait for the other client to send the encrypted private key.
	WaitPrivateKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WaitPrivateKeyResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Authorization(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error) {
	out := new(AuthorizationResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/Authorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetPublicKey(ctx context.Context, in *SetPublicKeyRequest, opts ...grpc.CallOption) (*SetPublicKeyResponse, error) {
	out := new(SetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/SetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error) {
	out := new(GetPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetDH(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDHResponse, error) {
	out := new(GetDHResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/GetDH", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendPrivateKey(ctx context.Context, in *SendPrivateKeyRequest, opts ...grpc.CallOption) (*SendPrivateKeyResponse, error) {
	out := new(SendPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/SendPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) WaitPrivateKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WaitPrivateKeyResponse, error) {
	out := new(WaitPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.auth.proto.Auth/WaitPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// Verify authentication Login with a valid user.
	// After successful authentication. The client must be registered to the device table for functions such as TOKEN checksum, managing devices, etc.
	// So you must submit a UA identifier in addition to the username and password when logging in.
	// A valid Token is returned and must be carried in subsequent API access operations.
	// https://datatracker.ietf.org/doc/html/rfc9068
	Authorization(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error)
	// SetPublicKey unlike activitypub, his private key is inaccessible to the server.
	// The public key is used for hvxahv privacy-related asymmetric encryption key.
	SetPublicKey(context.Context, *SetPublicKeyRequest) (*SetPublicKeyResponse, error)
	// GetPublicKey Use the account id to get the account public key.
	// Not activitypub public key.
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	// GetPrivateKey First, initiate a request to obtain the private key from the logged-in client.
	GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error)
	// GetDH Get the dh parameter.
	GetDH(context.Context, *emptypb.Empty) (*GetDHResponse, error)
	// SendPrivateKey Send the encrypted private key.
	SendPrivateKey(context.Context, *SendPrivateKeyRequest) (*SendPrivateKeyResponse, error)
	// WaitPrivateKey Wait for the other client to send the encrypted private key.
	WaitPrivateKey(context.Context, *emptypb.Empty) (*WaitPrivateKeyResponse, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Authorization(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
}
func (UnimplementedAuthServer) SetPublicKey(context.Context, *SetPublicKeyRequest) (*SetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPublicKey not implemented")
}
func (UnimplementedAuthServer) GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedAuthServer) GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedAuthServer) GetDH(context.Context, *emptypb.Empty) (*GetDHResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDH not implemented")
}
func (UnimplementedAuthServer) SendPrivateKey(context.Context, *SendPrivateKeyRequest) (*SendPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPrivateKey not implemented")
}
func (UnimplementedAuthServer) WaitPrivateKey(context.Context, *emptypb.Empty) (*WaitPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitPrivateKey not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Authorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/Authorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authorization(ctx, req.(*AuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/SetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetPublicKey(ctx, req.(*SetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetDH_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetDH(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/GetDH",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetDH(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/SendPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendPrivateKey(ctx, req.(*SendPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_WaitPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).WaitPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.auth.proto.Auth/WaitPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).WaitPrivateKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.auth.proto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorization",
			Handler:    _Auth_Authorization_Handler,
		},
		{
			MethodName: "SetPublicKey",
			Handler:    _Auth_SetPublicKey_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Auth_GetPublicKey_Handler,
		},
		{
			MethodName: "GetPrivateKey",
			Handler:    _Auth_GetPrivateKey_Handler,
		},
		{
			MethodName: "GetDH",
			Handler:    _Auth_GetDH_Handler,
		},
		{
			MethodName: "SendPrivateKey",
			Handler:    _Auth_SendPrivateKey_Handler,
		},
		{
			MethodName: "WaitPrivateKey",
			Handler:    _Auth_WaitPrivateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/auth/auth.proto",
}
