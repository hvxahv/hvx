/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"strconv"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	domain := viper.GetString("domain")
	b := NewActorsIsExist(domain, in.PreferredUsername).IsExist()
	return &pb.IsExistResponse{IsExist: b}, nil
}

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

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	actor, err := NewActors(in.PreferredUsername, in.PublicKey, in.ActorType).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Code: "200", ActorId: strconv.Itoa(int(actor.ID))}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	id, err := strconv.Atoi(in.GetActorId())
	if err != nil {
		return nil, err
	}
	actor, err := NewActorsId(uint(id)).Get()
	if err != nil {
		return nil, err
	}
	data := &pb.ActorData{
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
	}
	return &pb.GetResponse{
		Actor: data,
	}, nil
}

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

func (s *server) Edit(ctx context.Context, in *pb.EditRequest) (*pb.EditResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	a := new(Actors)
	a.ID = parse.ActorId
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

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := strconv.Atoi(in.GetId())
	if err != nil {
		return nil, err
	}
	if err := NewActorsId(uint(id)).Delete(); err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
