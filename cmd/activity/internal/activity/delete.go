package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cmd/activity/internal/outbox"
	"time"
)

type Delete struct {
	ActivityId string `json:"activityId"`
}

func (h *Handler) Delete(data []byte, to []string) (*pb.ActivityResponse, error) {
	d := Delete{}
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	// TODO - CREATE ARTICLE...
	var (
		notok []string
		ok    []string
		id    = fmt.Sprintf("%s/articles/%s/activity#create", h.aAddr, uuid.NewString())
		url   = fmt.Sprintf("%s/articles/%s/activity#create", h.aAddr, uuid.NewString())
	)

	var body = &activitypub.Delete{
		Context: activitypub.NewContext(),
		Id:      id,
		Type:    "Delete",
		Actor:   h.aAddr,
		To:      []string{"https://www.w3.org/ns/activitystreams#Public"},
		Object: struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			AtomUri string `json:"atomUri"`
		}{
			Id:      d.ActivityId,
			Type:    "Tombstone",
			AtomUri: url,
		},
		Signature: struct {
			Type           string    `json:"type"`
			Creator        string    `json:"creator"`
			Created        time.Time `json:"created"`
			SignatureValue string    `json:"signatureValue"`
		}{},
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// DELIVERY ...
	do, err := NewDelivery(marshal, h.aAddr, h.privateKey).Do(fmt.Sprintf("%s", h.inbox))
	if err != nil {
		return nil, err
	}
	if do.StatusCode != 202 {
		notok = append(notok, h.inbox)
		return nil, nil
	}
	ok = append(ok, h.inbox)
	if err := outbox.NewOutboxesDeleteByActivityId(h.actorId, d.ActivityId).Delete(); err != nil {
		return nil, err
	}
	return response(notok, ok)
}
