package internal

import (
	"context"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/saved"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *server) GetSaved(ctx context.Context, in *pb.GetSavedRequest) (*pb.Save, error) {
	return &pb.Save{}, nil
}

func (s *server) GetSaves(ctx context.Context, in *pb.GetSavesRequest) (*pb.GetSavesResponse, error) {
	return &pb.GetSavesResponse{}, nil
}

func (s *server) EditSaved(ctx context.Context, in *pb.EditSavedRequest) (*pb.EditSavedResponse, error) {
	return &pb.EditSavedResponse{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

func (s *server) DeleteSaves(ctx context.Context, in *pb.DeleteSavesRequest) (*pb.DeleteSavesResponse, error) {
	return &pb.DeleteSavesResponse{}, nil
}
