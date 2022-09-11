/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

func (s *server) Activity(ctx context.Context, in *pb.ActivityRequest) (*pb.ActivityResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	actor, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(strconv.Itoa(int(parse.ActorId)))
	if err != nil {
		return nil, err
	}
	// Actor address
	aAddr := actor.Actor.GetAddress()

	accounts, err := clientv1.New(ctx, microsvc.AccountServiceName).GetPrivateKey(strconv.Itoa(int(parse.AccountId)))
	if err != nil {
		return nil, err
	}
	privateKey := accounts.GetPrivateKey()

	var r *pb.ActivityResponse
	switch in.GetType() {
	case activitypub.FollowType:
		inbox := in.TO[0]
		follow, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Follow()
		if err != nil {
			return nil, err
		}
		r = follow
	case activitypub.AcceptType:
		inbox := in.TO[0]
		accept, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Accept([]byte(in.GetBody()))
		if err != nil {
			return nil, err
		}
		r = accept
	case activitypub.RejectType:
		inbox := in.TO[0]
		reject, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Reject([]byte(in.GetBody()))
		if err != nil {
			return nil, err
		}
		r = reject
	case activitypub.UndoType:
		inbox := in.TO[0]
		reject, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Undo([]byte(in.GetBody()))
		if err != nil {
			return nil, err
		}
		r = reject
	default:

	}
	return r, nil
}
