package internal

import (
	"context"
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/saved"
	"github.com/hvxahv/hvx/microsvc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewSaves(parse.AccountId, in.Name, in.Comment, in.Cid, in.FileType, in.IsPrivate).Create(); err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) GetSaved(ctx context.Context, in *pb.GetSavedRequest) (*pb.Save, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	saved, err := NewSavesId(uint(id)).GetSaved()
	if err != nil {
		return nil, err
	}
	return &pb.Save{
		Id:        strconv.Itoa(int(saved.ID)),
		Name:      saved.Name,
		Comment:   saved.Comment,
		Cid:       saved.Cid,
		Types:     saved.Types,
		CreatedAt: saved.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *server) GetSaves(ctx context.Context, in *emptypb.Empty) (*pb.GetSavesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	saves, err := NewSavesAccountId(parse.AccountId).GetSaves()
	if err != nil {
		return nil, err
	}

	var ret []*pb.Save
	for _, saved := range saves {
		ret = append(ret, &pb.Save{
			Id:        strconv.Itoa(int(saved.ID)),
			Name:      saved.Name,
			Comment:   saved.Comment,
			Cid:       saved.Cid,
			Types:     saved.Types,
			CreatedAt: saved.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &pb.GetSavesResponse{
		Code:  "200",
		Saves: ret,
	}, nil
}

func (s *server) EditSaved(ctx context.Context, in *pb.EditSavedRequest) (*pb.EditSavedResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	saved := new(Saves)
	switch {
	case in.Name != "":
		saved.Name = in.Name
	case in.Comment != "":
		saved.Comment = in.Comment
	}

	if err := saved.EditSaved(uint(id), parse.AccountId); err != nil {
		return nil, err
	}
	return &pb.EditSavedResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	if err := NewSavesDelete(uint(id), parse.AccountId).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) DeleteSaves(ctx context.Context, in *pb.DeleteSavesRequest) (*pb.DeleteSavesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewSavesAccountId(parse.AccountId).DeleteSaves(); err != nil {
		return nil, err
	}
	return &pb.DeleteSavesResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}
