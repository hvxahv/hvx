/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"encoding/json"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/mailer"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"net/mail"
	"net/url"
	"strconv"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	domain := viper.GetString("domain")
	b, ok := NewActorsIsExist(domain, in.PreferredUsername).IsExist()

	if !ok {
		return &pb.IsExistResponse{
			IsExist:   !ok,
			ActorType: "",
		}, nil
	}
	return &pb.IsExistResponse{
		IsExist:   ok,
		ActorType: b.ActorType,
	}, nil
}

func (s *server) IsRemoteExist(ctx context.Context, in *pb.IsRemoteExistRequest) (*pb.IsExistResponse, error) {
	b, exist := NewActorsIsExist(in.Domain, in.PreferredUsername).IsExist()
	if exist {
		return &pb.IsExistResponse{
			IsExist:   true,
			ActorType: b.ActorType,
		}, nil
	}
	return &pb.IsExistResponse{
		IsExist:   false,
		ActorType: "",
	}, nil
}

func (s *server) GetActorByUsername(ctx context.Context, in *pb.GetActorByUsernameRequest) (*pb.ActorData, error) {
	actor, err := NewAccountUsername(in.Username).GetActorByUsername()
	if err != nil {
		return nil, err
	}

	return &pb.ActorData{
		Id:                int64(actor.ID),
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
	switch in.ActorType {
	case "Channel":
		actor, err := NewChannels(in.PreferredUsername, in.PublicKey, in.ActorType).Create()
		if err != nil {
			return nil, err
		}
		return &pb.CreateResponse{Code: "200", ActorId: int64(actor.ID)}, nil
	case "Person":
		actor, err := NewActors(in.PreferredUsername, in.PublicKey, in.ActorType).Create()
		if err != nil {
			return nil, err
		}
		return &pb.CreateResponse{Code: "200", ActorId: int64(actor.ID)}, nil
	}
	return &pb.CreateResponse{Code: "500", ActorId: 0}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {

	actor, err := NewActorsId(uint(in.GetActorId())).Get()
	if err != nil {
		return nil, err
	}
	data := &pb.ActorData{
		Id:                int64(actor.ID),
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

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	var a []*pb.ActorData

	addr, err := mail.ParseAddress(in.GetPreferredUsername())
	if err == nil {
		format, err := mailer.ParseEmailAddress(in.GetPreferredUsername())
		if err != nil {
			return nil, err
		}

		exist, err := clientv1.New(ctx, microsvc.ActorServiceName).IsRemoteExist(format.Username, format.Domain)
		if err != nil {
			return nil, err
		}

		if exist.IsExist {
			v, err := NewPreferredUsernameAndDomain(format.Username, format.Domain).GetActorByUsername()
			if err != nil {
				return nil, err
			}
			var ad pb.ActorData
			ad.Id = int64(v.ID)
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

			return &pb.SearchResponse{Code: "200", Actors: a}, nil
		}
		handler, err := activitypub.GetWebFingerHandler(addr.Address)
		if err != nil {
			return nil, err
		}
		g, err := activitypub.GetActorByWebfinger(handler)
		if err != nil {
			return nil, err
		}

		var act activitypub.Actor
		if err := json.Unmarshal(g.Body(), &act); err != nil {
			return nil, err
		}
		parse, err := url.Parse(act.Url)
		if err != nil {
			return nil, err
		}
		actor, err := NewAddActors(act.PreferredUsername, parse.Host, act.Icon.Url, act.Name, act.Summary,
			act.Inbox, act.Url, act.PublicKey.PublicKeyPem, act.Type).AddActor()
		if err != nil {
			return nil, err
		}
		a = append(a, &pb.ActorData{
			Id:                int64(actor.ID),
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
		})
	}

	// Direct user search
	actors, err := NewPreferredUsername(in.GetPreferredUsername()).GetActorsByPreferredUsername()
	if err != nil {
		return nil, err
	}

	for _, v := range actors {
		var ad pb.ActorData
		ad.Id = int64(v.ID)
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

	return &pb.SearchResponse{Code: "200", Actors: a}, nil
}

func (s *server) GetActorByAddress(ctx context.Context, in *pb.GetActorByAddressRequest) (*pb.ActorData, error) {
	actor, err := NewActorAddress(in.GetAddress()).GetActorByAddress()
	if err != nil {
		return nil, err
	}

	return &pb.ActorData{
		Id:                int64(actor.ID),
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
	return &pb.EditResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if err := NewActorsId(uint(in.GetActorId())).Delete(); err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
