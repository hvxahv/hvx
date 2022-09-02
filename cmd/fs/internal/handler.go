package internal

import (
	"context"
	"github.com/hvxahv/hvx/APIs/v1alpha1/fs"
	"strconv"
)

func (f *server) Create(ctx context.Context, in *fs.CreateRequest) (*fs.CreateResponse, error) {
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	if err := NewFsCreate(uint(aid), in.FileName, in.Address).Create(); err != nil {
		return nil, err
	}
	return &fs.CreateResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (f *server) Delete(ctx context.Context, in *fs.DeleteRequest) (*fs.DeleteResponse, error) {
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	if err := NewFs(uint(aid), in.FileName).Delete(); err != nil {
		return nil, err
	}
	return &fs.DeleteResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (f *server) Get(ctx context.Context, in *fs.GetRequest) (*fs.GetResponse, error) {
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	fd, err := NewFs(uint(aid), in.FileName).Get()
	if err != nil {
		return nil, err
	}
	return &fs.GetResponse{
		Code:     "200",
		FileName: fd.FileName,
		Address:  fd.Address,
	}, nil
}
