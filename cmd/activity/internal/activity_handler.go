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
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func (s *server) Activity(ctx context.Context, in *pb.ActivityRequest) (*pb.ActivityResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	h, err := activity.NewHandler(
		*activity.NewObject(in.GetTO()),
		*activity.NewActor(int64(parse.ActorId), int64(parse.AccountId)),
	)
	if err != nil || h.Err != nil {
		return nil, errors.Newf(h.Err.Error(), err)
	}

	switch in.GetType() {
	case activitypub.FollowType:
		fmt.Println("FOLLOW")
		follow, err := h.Follow()
		if err != nil {
			return nil, err
		}
		return follow, nil

	case activitypub.AcceptType:
		accept, err := h.Accept([]byte(in.Body))
		if err != nil {
			return nil, err
		}
		return accept, nil

	case activitypub.RejectType:
		reject, err := h.Reject([]byte(in.Body))
		if err != nil {
			return nil, err
		}
		return reject, nil

	case activitypub.UndoType:
		undo, err := h.Undo([]byte(in.Body))
		if err != nil {
			return nil, err
		}
		return undo, nil

	default:

	}

	return &pb.ActivityResponse{
		Code:   "500",
		Status: "I_DON'T_KNOW_WHAT_TO_DO",
	}, nil
}

func (s *server) ArticleCreateActivity(ctx context.Context, in *pb.ArticleCreateActivityRequest) (*pb.ActivityResponse, error) {
	h := activity.Handler{Actor: *activity.NewActor(in.GetActorId(), in.GetAccountId())}
	var (
		failures  []string
		successes []string
	)

	if len(in.Article.GetTo()) > 0 {
		for _, address := range in.Article.GetTo() {
			h.Object = *activity.NewObject(address)
			create, err := h.Create(address, in)
			if err != nil {
				return nil, err
			}

			if create.Status == "failures" {
				failures = append(failures, create.Address)
			} else {
				successes = append(successes, create.Address)
			}
		}
	}
	if len(in.Article.GetTo()) < 1 {
		// Send to followers.
		in.Article.To = append(in.Article.To, fmt.Sprintf("%s/followers", h.Actor.Address))

		get, err := friendship.NewFollows(uint(in.GetActorId()), friendship.Follower).Get()
		if err != nil {
			return nil, err
		}

		for _, i := range get {
			h.Object = *activity.NewObjectId(i)
			create, err := h.Create(h.Object.Address, in)
			if err != nil {
				return nil, err
			}
			if create.Status == "failures" {
				failures = append(failures, create.Address)
			} else {
				successes = append(successes, create.Address)
			}
		}

	}
	if len(in.Article.GetAudience()) > 0 {
		// SYNC TO CHANNELS
		// GET CHANNELS SUB
		//
		for _, i := range in.Article.GetAudience() {
			in.Article.Cc = append(in.Article.Cc, fmt.Sprintf("https://%s/c/%s", viper.GetString("domain"), strconv.Itoa(int(i))))
			subscribers, err := clientv1.New(ctx, microsvc.ChannelServiceName).GetSubscribers(i, in.GetActorId())
			if err != nil {
				return nil, err
			}
			for _, sub := range subscribers.Subscriber {
				h.Object = *activity.NewObjectId(uint(sub.GetId()))
				create, err := h.Create(h.Object.Address, in)
				if err != nil {
					return nil, err
				}
				if create.Status == "failures" {
					failures = append(failures, create.Address)
				} else {
					successes = append(successes, create.Address)
				}
			}
		}
	}
	fmt.Println(failures, successes)
	return &pb.ActivityResponse{
		Code:      "200",
		Status:    "ok",
		Successes: successes,
		Failures:  failures,
	}, nil

}
