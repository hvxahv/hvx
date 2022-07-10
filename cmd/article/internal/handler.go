package internal

import (
	"context"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/article"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}

func (s *server) GetArticles(ctx context.Context, in *pb.GetArticlesRequest) (*pb.GetArticlesResponse, error) {
	return &pb.GetArticlesResponse{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return &pb.UpdateResponse{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}
