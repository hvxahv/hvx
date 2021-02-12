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

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	NewArticle(ctx context.Context, in *ArticleData, opts ...grpc.CallOption) (*NewArticleReply, error)
	GetArticles(ctx context.Context, in *GetArticleData, opts ...grpc.CallOption) (*GetArticleReply, error)
	UpdateArticle(ctx context.Context, in *ArticleData, opts ...grpc.CallOption) (*UpdateArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleByID, opts ...grpc.CallOption) (*DeleteArticleReply, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) NewArticle(ctx context.Context, in *ArticleData, opts ...grpc.CallOption) (*NewArticleReply, error) {
	out := new(NewArticleReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Articles/NewArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) GetArticles(ctx context.Context, in *GetArticleData, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Articles/GetArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) UpdateArticle(ctx context.Context, in *ArticleData, opts ...grpc.CallOption) (*UpdateArticleReply, error) {
	out := new(UpdateArticleReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Articles/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) DeleteArticle(ctx context.Context, in *DeleteArticleByID, opts ...grpc.CallOption) (*DeleteArticleReply, error) {
	out := new(DeleteArticleReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1.proto.Articles/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	NewArticle(context.Context, *ArticleData) (*NewArticleReply, error)
	GetArticles(context.Context, *GetArticleData) (*GetArticleReply, error)
	UpdateArticle(context.Context, *ArticleData) (*UpdateArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleByID) (*DeleteArticleReply, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) NewArticle(context.Context, *ArticleData) (*NewArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewArticle not implemented")
}
func (UnimplementedArticlesServer) GetArticles(context.Context, *GetArticleData) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedArticlesServer) UpdateArticle(context.Context, *ArticleData) (*UpdateArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticlesServer) DeleteArticle(context.Context, *DeleteArticleByID) (*DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticlesServer) mustEmbedUnimplementedArticlesServer() {}

// UnsafeArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServer will
// result in compilation errors.
type UnsafeArticlesServer interface {
	mustEmbedUnimplementedArticlesServer()
}

func RegisterArticlesServer(s grpc.ServiceRegistrar, srv ArticlesServer) {
	s.RegisterService(&Articles_ServiceDesc, srv)
}

func _Articles_NewArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).NewArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Articles/NewArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).NewArticle(ctx, req.(*ArticleData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Articles/GetArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).GetArticles(ctx, req.(*GetArticleData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Articles/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).UpdateArticle(ctx, req.(*ArticleData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1.proto.Articles/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).DeleteArticle(ctx, req.(*DeleteArticleByID))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1.proto.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewArticle",
			Handler:    _Articles_NewArticle_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _Articles_GetArticles_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Articles_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Articles_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hvxahv/v1/articles.proto",
}
