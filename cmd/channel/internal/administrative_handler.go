package internal

import (
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
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
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	addedID, err := strconv.Atoi(in.AddedId)
	if err != nil {
		return nil, err
	}
	if err := NewAdministratesAdd(uint(channelId), parse.ActorId).AddAdministrator(uint(addedID)); err != nil {
		return nil, err
	}

	return &pb.AddAdministratorResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) RemoveAdministrator(ctx context.Context, in *pb.RemoveAdministratorRequest) (*pb.RemoveAdministratorResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	removedID, err := strconv.Atoi(in.RemovedId)
	if err != nil {
		return nil, err
	}
	if err := NewAdministratesPermission(uint(channelId), parse.ActorId).RemoveAdministrator(uint(removedID)); err != nil {
		return nil, err
	}
	return &pb.RemoveAdministratorResponse{
		Code:  "200",
		Reply: "ok",
	}, nil
}

func (s *server) GetAdministrators(ctx context.Context, in *pb.GetAdministratorsRequest) (*pb.GetAdministratorsResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	administrators, err := NewAdministratesPermission(uint(channelId), parse.ActorId).GetAdministrators()
	if err != nil {
		return nil, err
	}

	var admins []*pb.AdminsData
	for _, a := range administrators {
		var admin pb.AdminsData

		actor, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(strconv.Itoa(int(a.AdminId)))
		if err != nil {
			return nil, err
		}
		admin.IsOwner = a.IsOwner
		admin.Admin = actor.Actor
		admins = append(admins, &admin)
	}

	return &pb.GetAdministratorsResponse{
		Code:   "200",
		Admins: admins,
	}, nil
}

func (s *server) ExitAdministrator(ctx context.Context, in *pb.ExitAdministratorRequest) (*pb.ExitAdministratorResponse, error) {
	return &pb.ExitAdministratorResponse{}, nil
}
