package internal

import (
	"context"
	"github.com/hvxahv/hvx/APIs/v1alpha1/fs"
)

func (f *server) Create(ctx context.Context, in *fs.CreateRequest) (*fs.CreateResponse, error) {
	if err := NewFsCreate(uint(in.GetAccountId()), in.FileName, in.Address).Create(); err != nil {
		return nil, err
	}
	return &fs.CreateResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (f *server) Delete(ctx context.Context, in *fs.DeleteRequest) (*fs.DeleteResponse, error) {
	if err := NewFs(uint(in.GetAccountId()), in.FileName).Delete(); err != nil {
		return nil, err
	}
	return &fs.DeleteResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (f *server) Get(ctx context.Context, in *fs.GetRequest) (*fs.GetResponse, error) {
	fd, err := NewFs(uint(in.GetAccountId()), in.FileName).Get()
	if err != nil {
		return nil, err
	}
	return &fs.GetResponse{
		Code:     "200",
		FileName: fd.FileName,
		Address:  fd.Address,
	}, nil
}
