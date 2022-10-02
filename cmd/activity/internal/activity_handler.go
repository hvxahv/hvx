/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"golang.org/x/net/context"
)

func (s *server) Activity(ctx context.Context, in *pb.ActivityRequest) (*pb.ActivityResponse, error) {

	fmt.Println(in)
	//parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//actors, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(parse.ActorId))
	//if err != nil {
	//	return nil, err
	//}
	//// Actor address
	//aAddr := actors.Actor.GetAddress()
	//
	//accounts, err := clientv1.New(ctx, microsvc.AccountServiceName).GetPrivateKey(int64(parse.AccountId))
	//if err != nil {
	//	return nil, err
	//}
	//privateKey := accounts.GetPrivateKey()
	//
	//var r *pb.ActivityResponse
	//switch in.GetType() {
	//case activitypub.FollowType:
	//	inbox := in.Delivery.To[0]
	//	follow, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Follow()
	//	if err != nil {
	//		return nil, err
	//	}
	//	r = follow
	//
	//case activitypub.AcceptType:
	//	inbox := in.TO[0]
	//	accept, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Accept([]byte(in.GetBody()))
	//	if err != nil {
	//		return nil, err
	//	}
	//	r = accept
	//
	//case activitypub.RejectType:
	//	inbox := in.TO[0]
	//	reject, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Reject([]byte(in.GetBody()))
	//	if err != nil {
	//		return nil, err
	//	}
	//	r = reject
	//
	//case activitypub.UndoType:
	//	inbox := in.TO[0]
	//	reject, err := activity.NewHandler(inbox, aAddr, privateKey, parse.ActorId).Undo([]byte(in.GetBody()))
	//	if err != nil {
	//		return nil, err
	//	}
	//	r = reject
	//
	//// COMPLEX ACTIVITY HANDLING....
	//case activitypub.CreateType:
	//	//inbox := in.TO[0]
	//	// TODO - FIX RANGE TO CC BTO BCC...
	//	var followers []string
	//	var ad []*actor.ActorData
	//	if len(in.TO) == 0 {
	//		// GET FOLLOWER
	//		f, err := friendship.NewFollows(parse.ActorId, friendship.Follower).Get()
	//		if err != nil {
	//			return nil, err
	//		}
	//		for _, i := range f {
	//			a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(i))
	//			if err != nil {
	//				return nil, err
	//			}
	//			ad = append(ad, a.Actor)
	//			followers = append(followers, a.Actor.GetAddress())
	//		}
	//	}
	//	if len(in.GetSync()) != 0 {
	//		for _, i := range in.GetSync() {
	//			subscribers, err := clientv1.New(ctx, "channel").GetSubscribers(i)
	//			if err != nil {
	//				return nil, err
	//			}
	//			fmt.Println(subscribers)
	//		}
	//	}
	//
	//	for _, i := range ad {
	//		_, err := activity.NewHandler(i.Inbox,
	//			aAddr,
	//			privateKey,
	//			parse.ActorId,
	//		).
	//			Create(
	//				[]byte(in.GetBody()),
	//				followers,
	//				in.CC,
	//			)
	//		if err != nil {
	//			r.Status = "ERROR"
	//			r.Code = "501"
	//		}
	//	}
	//
	//	// TODO - FIX RESPONSE ...
	//	return &pb.ActivityResponse{
	//		Code:   "200",
	//		Status: "ok",
	//		Inbox:  nil,
	//	}, nil
	//
	//case activitypub.DeleteType:
	//	var followers []string
	//	var ad []*actor.ActorData
	//	if len(in.TO) == 0 {
	//		// GET FOLLOWER
	//		f, err := friendship.NewFollows(parse.ActorId, friendship.Follower).Get()
	//		if err != nil {
	//			return nil, err
	//		}
	//		for _, i := range f {
	//			a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(i))
	//			if err != nil {
	//				return nil, err
	//			}
	//			ad = append(ad, a.Actor)
	//			followers = append(followers, a.Actor.GetAddress())
	//		}
	//	}
	//
	//	for _, i := range ad {
	//		_, err := activity.NewHandler(i.Inbox,
	//			aAddr,
	//			privateKey,
	//			parse.ActorId,
	//		).
	//			Delete(
	//				[]byte(in.GetBody()),
	//				followers,
	//			)
	//		if err != nil {
	//			r.Status = "ERROR"
	//			r.Code = "501"
	//		}
	//	}
	//
	//	// TODO - FIX RESPONSE ...
	//	return &pb.ActivityResponse{
	//		Code:   "200",
	//		Status: "ok",
	//		Inbox:  nil,
	//	}, nil
	//
	//default:
	//
	//}
	//return r, nil

	return nil, nil
}
