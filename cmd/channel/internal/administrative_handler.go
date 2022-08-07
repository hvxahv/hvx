package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"golang.org/x/net/context"
)

// Administrative ...

func (s *server) IsAdministrator(ctx context.Context, in *pb.IsAdministratorRequest) (*pb.IsAdministratorResponse, error) {
	return &pb.IsAdministratorResponse{}, nil
}

func (s *server) IsOwner(ctx context.Context, in *pb.IsOwnerRequest) (*pb.IsOwnerResponse, error) {
	return &pb.IsOwnerResponse{}, nil
}

func (s *server) AddAdministrator(ctx context.Context, in *pb.AddAdministratorRequest) (*pb.AddAdministratorResponse, error) {
	return &pb.AddAdministratorResponse{}, nil
}

func (s *server) RemoveAdministrator(ctx context.Context, in *pb.RemoveAdministratorRequest) (*pb.RemoveAdministratorResponse, error) {
	return &pb.RemoveAdministratorResponse{}, nil
}

func (s *server) GetAdministrators(ctx context.Context, in *pb.GetAdministratorsRequest) (*pb.GetAdministratorsResponse, error) {
	return &pb.GetAdministratorsResponse{}, nil
}

func (s *server) ExitAdministrator(ctx context.Context, in *pb.ExitAdministratorRequest) (*pb.ExitAdministratorResponse, error) {
	return &pb.ExitAdministratorResponse{}, nil
}
