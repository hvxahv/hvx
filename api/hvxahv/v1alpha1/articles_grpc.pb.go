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

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	NewArticles(ctx context.Context, in *NewArticlesData, opts ...grpc.CallOption) (*ArticlesReply, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) NewArticles(ctx context.Context, in *NewArticlesData, opts ...grpc.CallOption) (*ArticlesReply, error) {
	out := new(ArticlesReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Articles/NewArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	NewArticles(context.Context, *NewArticlesData) (*ArticlesReply, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) NewArticles(context.Context, *NewArticlesData) (*ArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewArticles not implemented")
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

func _Articles_NewArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewArticlesData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).NewArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Articles/NewArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).NewArticles(ctx, req.(*NewArticlesData))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewArticles",
			Handler:    _Articles_NewArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hvxahv/v1alpha1/articles.proto",
}
