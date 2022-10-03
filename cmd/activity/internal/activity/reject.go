package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cmd/activity/internal/delivery"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
)

type rejectObject struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

func (h *Handler) Reject(data []byte) (*pb.ActivityResponse, error) {
	var (
		failures  []string
		successes []string
	)
	var reject rejectObject
	if err := json.Unmarshal(data, &reject); err != nil {
		return nil, err
	}
	var (
		id = fmt.Sprintf("%s/#rejects/%s", h.Actor.Address, uuid.NewString())
	)

	body := &activitypub.Reject{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      id,
		Type:    activitypub.RejectType,
		Actor:   h.Actor.Address,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     reject.Id,
			Type:   reject.Type,
			Actor:  reject.Actor,
			Object: reject.Object,
		},
	}
	marshal, err := json.Marshal(body)
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

	// CREATE FOLLOW OUTBOX ...
	if err := outbox.NewOutboxes(uint(h.Actor.Id), h.ActivityId, h.Object.Address, activitypub.RejectType, string(marshal)).Create(); err != nil {
		return nil, err
	}

	return &pb.ActivityResponse{
		Code:      "200",
		Status:    "ok",
		Successes: successes,
		Failures:  failures,
	}, nil
}
