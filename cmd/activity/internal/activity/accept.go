package activity

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/delivery"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
	"github.com/hvxahv/hvx/microsvc"
)

type acceptObject struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

func (h *Handler) Accept(data []byte) (*pb.ActivityResponse, error) {
	var (
		failures  []string
		successes []string
	)
	var accept acceptObject
	if err := json.Unmarshal(data, &accept); err != nil {
		return nil, err
	}
	var (
		id = fmt.Sprintf("%s/#rejects/%s", h.Actor.Address, uuid.NewString())
	)

	//MARSHAL DATA.
	marshal, err := json.Marshal(&activitypub.Accept{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      id,
		Type:    activitypub.AcceptType,
		Actor:   h.Actor.Address,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     accept.Id,
			Type:   accept.Type,
			Actor:  accept.Actor,
			Object: accept.Object,
		},
	})
	if err != nil {
		return nil, err
	}

	// DELIVERY ...
	do, err := delivery.New(h.Actor.PublicKeyId, h.Actor.PrivateKey, marshal).Do(h.Object.Inbox)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != 202 {
		failures = append(failures, h.Object.Address)
		return nil, nil
	}
	successes = append(successes, h.Object.Address)

	// CREATE FOLLOW OUTBOX ...
	if err := outbox.NewOutboxes(uint(h.Actor.Id), h.ActivityId, h.Object.Address, activitypub.AcceptType, string(marshal)).Create(); err != nil {
		return nil, err
	}

	switch accept.Type {
	case activitypub.FollowType:
		object, err := clientv1.New(context.Background(), microsvc.ActorServiceName).GetActorByAddress(h.Object.Address)
		if err != nil {
			return nil, err
		}

		// IF ACCEPT FOLLOW REQUEST, ADD FOLLOWER.
		if err := friendship.NewFollower(uint(h.Actor.Id), uint(object.Id)).Follow(); err != nil {
			return nil, err
		}
	default:
	}

	return &pb.ActivityResponse{
		Code:      "200",
		Status:    "ok",
		Successes: successes,
		Failures:  failures,
	}, nil
}
