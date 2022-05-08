/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package device

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	"github.com/hvxahv/hvx/pkg/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *server) DeleteDeviceAllByAccountID(ctx context.Context, g *emptypb.Empty) (*pb.DeleteDeviceAllByAccountIDResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewDevicesAccountID(id).DeleteAllByAccountID(); err != nil {
		return nil, err
	}
	return &pb.DeleteDeviceAllByAccountIDResponse{Code: "200", Reply: "ok"}, nil
}
