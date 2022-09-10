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
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

const (
	Follow = "Follow"
	Accept = "Accept"
	Reject = "Reject"
	Undo   = "Undo"
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
	defer actorc.Close()
	actors, err := actor.NewActorClient(actorc.Conn).Get(ctx, &actor.GetRequest{
		ActorId: strconv.Itoa(int(parse.ActorId)),
	})
	if err != nil {
		return nil, err
	}
	actorAddress := actors.Actor.GetAddress()

	acctc, err := clientv1.New(ctx, microsvc.NewGRPCAddress("account").Get())
	if err != nil {
		return nil, err
	}
	defer acctc.Close()
	accounts, err := account.NewAccountsClient(acctc.Conn).GetPrivateKey(ctx, &account.GetPrivateKeyRequest{
		AccountId: strconv.Itoa(int(parse.AccountId)),
	})
	if err != nil {
		return nil, err
	}
	privateKey := accounts.GetPrivateKey()

	var notok []string
	var ok []string
	switch in.GetType() {
	case Follow:
		var id = fmt.Sprintf("%s/%s", actorAddress, uuid.NewString())
		inbox := in.TO[0]
		body := &activitypub.Follow{
			Context: "https://www.w3.org/ns/activitystreams",
			Id:      id,
			Type:    "Follow",
			Actor:   actorAddress,
			Object:  inbox,
		}
		marshal, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		instances, err := activity.IsLocalInstance(inbox)
		if err != nil {
			return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
		}
		if !instances {
			do, err := activity.NewSender(marshal, actorAddress, privateKey).Do(fmt.Sprintf("%s/inbox", inbox))
			if err != nil {
				return nil, err
			}
			if do.StatusCode != 202 {
				notok = append(notok, inbox)
				return nil, nil
			}
			ok = append(ok, inbox)
		}
		// TODO - SEND TO LOCAL USER INBOX...

		to := outbox.Stos(in.GetTO())
		if err := outbox.NewFollowOutboxes(parse.ActorId, id, to, Follow, string(marshal)).Create(); err != nil {
			return nil, err
		}

	case Accept:
		var b activity.Body
		if err := json.Unmarshal([]byte(in.GetBody()), &b); err != nil {
			return nil, err
		}
		var id = fmt.Sprintf("%s/#accept/%s", actorAddress, uuid.NewString())
		inbox := in.TO[0]
		body := &activitypub.Accept{
			Context: "https://www.w3.org/ns/activitystreams",
			Id:      id,
			Type:    in.GetType(),
			Actor:   actorAddress,
			Object: struct {
				Id     string `json:"id"`
				Type   string `json:"type"`
				Actor  string `json:"actor"`
				Object string `json:"object"`
			}{
				Id:     b.Id,
				Type:   b.Type,
				Actor:  b.Actor,
				Object: b.Object,
			},
		}
		marshal, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		instances, err := activity.IsLocalInstance(inbox)
		if err != nil {
			return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
		}
		if !instances {
			do, err := activity.NewSender(marshal, actorAddress, privateKey).Do(fmt.Sprintf("%s/inbox", inbox))
			if err != nil {
				return nil, err
			}
			if do.StatusCode != 202 {
				notok = append(notok, inbox)
				return nil, nil
			}
			ok = append(ok, inbox)
		}

		to := outbox.Stos(in.GetTO())
		if err := outbox.NewFollowOutboxes(parse.ActorId, id, to, Accept, string(marshal)).Create(); err != nil {
			return nil, err
		}
		switch b.Type {
		case Follow:
			//If the accept is of type follow
			// TODO - CREATE FOLLOW TABLE FOLLOWER...
			gba, err := actor.NewActorClient(actorc.Conn).GetActorByAddress(ctx, &actor.GetActorByAddressRequest{
				Address: inbox,
			})
			if err != nil {
				return nil, err
			}

			actorId, err := strconv.Atoi(actors.Actor.Id)
			if err != nil {
				return nil, err
			}
			targetId, err := strconv.Atoi(gba.Id)
			if err != nil {
				return nil, err
			}
			if err := friendship.NewFollower(uint(actorId), uint(targetId)).Follow(); err != nil {
				return nil, err
			}
		default:

		}
	case Reject:
		var b activity.Body
		if err := json.Unmarshal([]byte(in.GetBody()), &b); err != nil {
			return nil, err
		}
		inbox := in.TO[0]
		body := &activitypub.Reject{
			Context: "https://www.w3.org/ns/activitystreams",
			Id:      fmt.Sprintf("%s/#rejects/%s", actorAddress, uuid.NewString()),
			Type:    in.GetType(),
			Actor:   actorAddress,
			Object: struct {
				Id     string `json:"id"`
				Type   string `json:"type"`
				Actor  string `json:"actor"`
				Object string `json:"object"`
			}{
				Id:     b.Id,
				Type:   b.Type,
				Actor:  b.Actor,
				Object: b.Object,
			},
		}
		marshal, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		instances, err := activity.IsLocalInstance(inbox)
		if err != nil {
			return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
		}
		if !instances {
			do, err := activity.NewSender(marshal, actorAddress, privateKey).Do(fmt.Sprintf("%s/inbox", inbox))
			if err != nil {
				return nil, err
			}
			if do.StatusCode != 202 {
				notok = append(notok, inbox)
				return nil, nil
			}
			ok = append(ok, inbox)
		}
		// TODO - DELETE INBOX FOLLOW REQUEST...
	case Undo:
		var b activity.Body
		if err := json.Unmarshal([]byte(in.GetBody()), &b); err != nil {
			return nil, err
		}
		inbox := in.TO[0]
		body := &activitypub.Reject{
			Context: "https://www.w3.org/ns/activitystreams",
			Id:      fmt.Sprintf("%s/#undo/%s", actorAddress, uuid.NewString()),
			Type:    in.GetType(),
			Actor:   actorAddress,
			Object: struct {
				Id     string `json:"id"`
				Type   string `json:"type"`
				Actor  string `json:"actor"`
				Object string `json:"object"`
			}{
				Id:     b.Id,
				Type:   b.Type,
				Actor:  b.Actor,
				Object: b.Object,
			},
		}
		marshal, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		instances, err := activity.IsLocalInstance(inbox)
		if err != nil {
			return nil, errors.New("INBOX_URL_FAILS_WHEN_PARSING")
		}
		if !instances {
			do, err := activity.NewSender(marshal, actorAddress, privateKey).Do(fmt.Sprintf("%s/inbox", inbox))
			if err != nil {
				return nil, err
			}
			if do.StatusCode != 202 {
				notok = append(notok, inbox)
				return nil, nil
			}
			ok = append(ok, inbox)
			if err := outbox.NewOutboxesActivityId(b.Id).Delete(); err != nil {
				return nil, err
			}
		}
		// TODO - SEND TO LOCAL USER INBOX...
	default:

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
