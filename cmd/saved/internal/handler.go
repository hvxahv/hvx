package internal

import (
	"context"
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/saved"
	"github.com/hvxahv/hvx/microsvc"
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
	saved := new(Saves)
	switch {
	case in.Name != "":
		saved.EditSavedName(in.Name)
	case in.Comment != "":
		saved.EditSavedComment(in.Comment)
	}
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	accountId, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := saved.EditSaved(uint(id), accountId); err != nil {
		return nil, err
	}

	return &pb.EditSavedResponse{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

func (s *server) DeleteSaves(ctx context.Context, in *pb.DeleteSavesRequest) (*pb.DeleteSavesResponse, error) {
	return &pb.DeleteSavesResponse{}, nil
}
