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

type undoObject struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}

func (h *Handler) Undo(data []byte) (*pb.ActivityResponse, error) {
	var (
		failures  []string
		successes []string
	)
	var undo undoObject
	if err := json.Unmarshal(data, &undo); err != nil {
		return nil, err
	}
	var (
		id = fmt.Sprintf("%s/#rejects/%s", h.Actor.Address, uuid.NewString())
	)

	//MARSHAL DATA.
	marshal, err := json.Marshal(&activitypub.Accept{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      id,
		Type:    activitypub.UndoType,
		Actor:   h.Actor.Address,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     undo.Id,
			Type:   undo.Type,
			Actor:  undo.Actor,
			Object: undo.Object,
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

	if err := outbox.NewOutboxesDeleteByActivityId(uint(h.Actor.Id), undo.Id).Delete(); err != nil {
		return nil, err
	}
	return &pb.ActivityResponse{
		Code:      "200",
		Status:    "ok",
		Successes: successes,
		Failures:  failures,
	}, nil

}
