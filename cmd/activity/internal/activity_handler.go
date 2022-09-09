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
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

func (s *server) Activity(ctx context.Context, in *pb.ActivityRequest) (*pb.ActivityResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	// GET ACTOR DATA
	actorc, err := clientv1.New(ctx, microsvc.NewGRPCAddress("actor").Get())
	if err != nil {
		return nil, err
	}
	g, err := actor.NewActorClient(actorc.Conn).Get(ctx, &actor.GetRequest{
		ActorId: strconv.Itoa(int(parse.ActorId)),
	})
	if err != nil {
		return nil, err
	}
	actorAddress := g.Actor.GetAddress()

	acct, err := clientv1.New(ctx, microsvc.NewGRPCAddress("account").Get())
	if err != nil {
		return nil, err
	}
	gpk, err := account.NewAccountsClient(acct.Conn).GetPrivateKey(ctx, &account.GetPrivateKeyRequest{
		AccountId: strconv.Itoa(int(parse.AccountId)),
	})
	if err != nil {
		return nil, err
	}
	privateKey := gpk.GetPrivateKey()

	var notok []string
	var ok []string
	switch in.GetEvent() {
	case "Follow":
		for _, i := range in.TO {
			body := &activitypub.Follow{
				Context: "https://www.w3.org/ns/activitystreams",
				Id:      fmt.Sprintf("%s/%s", actorAddress, uuid.NewString()),
				Type:    "Follow",
				Actor:   actorAddress,
				Object:  i,
			}
			marshal, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			instances, err := activity.IsLocalInstance(i)
			if err != nil {
				return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
			}
			if !instances {
				do, err := activity.NewSender(marshal, actorAddress, privateKey).Do(fmt.Sprintf("%s/inbox", i))
				if err != nil {
					return nil, err
				}
				if do.StatusCode != 202 {
					notok = append(notok, i)
				}
				ok = append(ok, i)
			}
			// TODO - SEND TO LOCAL USER INBOX...
		}
	}

	// TODO - Send results to outbox...

	if len(notok) != 0 {
		return &pb.ActivityResponse{
			Code:   "200",
			Status: "THERE_ARE_REQUESTS_THAT_FAILED_TO_SEND",
			Inbox:  notok,
		}, nil
	}
	return &pb.ActivityResponse{
		Code:   "200",
		Status: "SEND_SUCCESSFULLY",
		Inbox:  ok,
	}, nil
}
