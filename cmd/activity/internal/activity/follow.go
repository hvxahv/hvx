package activity

import (
	"encoding/json"
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cmd/activity/internal/delivery"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
)

func (h *Handler) Follow() (*pb.ActivityResponse, error) {
	var (
		failures  []string
		successes []string
	)

	//  MARSHAL DATA.
	body := &activitypub.Follow{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      h.ActivityId,
		Type:    activitypub.FollowType,
		Actor:   h.Actor.Address,
		Object:  h.Object.Address,
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// DELIVERY ...

	fmt.Println(h.Actor.PublicKeyId, h.Actor.PrivateKey, marshal, h.Object.Inbox)
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
	if err := outbox.NewOutboxes(uint(h.Actor.Id), h.ActivityId, h.Object.Address, activitypub.FollowType, string(marshal)).Create(); err != nil {
		return nil, err
	}

	return &pb.ActivityResponse{
		Code:      "200",
		Status:    "ok",
		Successes: successes,
		Failures:  failures,
	}, nil
}
