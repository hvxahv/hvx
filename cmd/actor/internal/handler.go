/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/actor"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

// GetActorByUsername ...
func (s *server) GetActorByUsername(ctx context.Context, in *pb.GetActorByUsernameRequest) (*pb.ActorData, error) {
	actor, err := NewAccountUsername(in.Username).GetActorByUsername()
	if err != nil {
		return nil, err
	}

	return &pb.ActorData{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

// Create actor.
func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	actor, err := NewActors(in.PreferredUsername, in.PublicKey, in.ActorType).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Code: "200", ActorId: strconv.Itoa(int(actor.ID))}, nil
}

// GetActorsByPreferredUsername ...
func (s *server) GetActorsByPreferredUsername(ctx context.Context, in *pb.GetActorsByPreferredUsernameRequest) (*pb.GetActorsByPreferredUsernameResponse, error) {
	actors, err := NewPreferredUsername(in.GetPreferredUsername()).GetActorsByPreferredUsername()
	if err != nil {
		return nil, err
	}

	var a []*pb.ActorData
	for _, v := range actors {
		var ad pb.ActorData
		ad.Id = strconv.Itoa(int(v.ID))
		ad.PreferredUsername = v.PreferredUsername
		ad.Domain = v.Domain
		ad.Avatar = v.Avatar
		ad.Name = v.Name
		ad.Summary = v.Summary
		ad.Inbox = v.Inbox
		ad.Address = v.Address
		ad.PublicKey = v.PublicKey
		ad.ActorType = v.ActorType
		ad.IsRemote = strconv.FormatBool(v.IsRemote)
		a = append(a, &ad)
	}

	return &pb.GetActorsByPreferredUsernameResponse{Code: "200", Actors: a}, nil
}

// GetActorByAddress ...
func (s *server) GetActorByAddress(ctx context.Context, in *pb.GetActorByAddressRequest) (*pb.ActorData, error) {
	actor, err := NewActorAddress(in.GetAddress()).GetActorByAddress()
	if err != nil {
		return nil, err
	}

	return &pb.ActorData{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

// Edit ...
func (s *server) Edit(ctx context.Context, in *pb.EditRequest) (*pb.EditResponse, error) {
	userdata, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	actorId, err := strconv.Atoi(userdata.ActorId)
	if err != nil {
		return nil, err
	}

	a := new(Actors)
	a.ID = uint(actorId)
	if in.Avatar != "" {
		a.Avatar = in.Avatar
	}
	if in.Name != "" {
		a.Name = in.Name
	}
	if in.Summary != "" {
		a.Summary = in.Summary
	}
	if err := a.Edit(); err != nil {
		return nil, err
	}
	return &pb.EditResponse{Code: "200", Reply: "ok"}, nil
}

// Delete ...
func (s *server) Delete(ctx context.Context, in *empty.Empty) (*pb.DeleteResponse, error) {
	//actorId, err := microsvc.GetActorIdByTokenWithContext(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//if err := NewActorId(actorId).Delete(); err != nil {
	//	return nil, err
	//}
	return nil, nil
}
