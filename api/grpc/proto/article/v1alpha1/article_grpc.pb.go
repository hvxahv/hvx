// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha

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

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	// GetArticle Get the content of an article or status by its article ID.
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error)
	// GetArticlesByAccountID Get all posts or statuses published under this account by account ID.
	GetArticlesByAccountID(ctx context.Context, in *GetArticlesByAccountIDRequest, opts ...grpc.CallOption) (*GetArticlesByAccountIDResponse, error)
	// CreateArticle Create article or status.
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	// UpdateArticle Edit article or status content.
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
	// DeleteArticle Delete a post or status by post or status ID.
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleResponse, error) {
	out := new(GetArticleResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Article/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticlesByAccountID(ctx context.Context, in *GetArticlesByAccountIDRequest, opts ...grpc.CallOption) (*GetArticlesByAccountIDResponse, error) {
	out := new(GetArticlesByAccountIDResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Article/GetArticlesByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Article/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	out := new(UpdateArticleResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Article/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Article/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations should embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	// GetArticle Get the content of an article or status by its article ID.
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error)
	// GetArticlesByAccountID Get all posts or statuses published under this account by account ID.
	GetArticlesByAccountID(context.Context, *GetArticlesByAccountIDRequest) (*GetArticlesByAccountIDResponse, error)
	// CreateArticle Create article or status.
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	// UpdateArticle Edit article or status content.
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
	// DeleteArticle Delete a post or status by post or status ID.
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
}

// UnimplementedArticleServer should be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServer) GetArticlesByAccountID(context.Context, *GetArticlesByAccountIDRequest) (*GetArticlesByAccountIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByAccountID not implemented")
}
func (UnimplementedArticleServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Article/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticlesByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticlesByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Article/GetArticlesByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticlesByAccountID(ctx, req.(*GetArticlesByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Article/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Article/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Article/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticle",
			Handler:    _Article_GetArticle_Handler,
		},
		{
			MethodName: "GetArticlesByAccountID",
			Handler:    _Article_GetArticlesByAccountID_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Article_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Article_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Article_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/article/v1alpha1/article.proto",
}