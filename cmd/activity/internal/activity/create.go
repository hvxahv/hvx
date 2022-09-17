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

type Create struct {
	Type       string        `json:"type"`
	Title      string        `json:"title"`
	Summary    string        `json:"summary"`
	InReplyTo  interface{}   `json:"inReplyTo"`
	Content    string        `json:"content"`
	Attachment []interface{} `json:"attachment"`
	Tag        []interface{} `json:"tag"`
}

func (h *Handler) Create(data []byte, to, cc []string) (*pb.ActivityResponse, error) {
	c := Create{}
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	// TODO - CREATE ARTICLE...
	var (
		notok []string
		ok    []string
		id    = fmt.Sprintf("%s/articles/%s/activity#create", h.aAddr, uuid.NewString())
		url   = fmt.Sprintf("%s/articles/%s/activity#create", h.aAddr, uuid.NewString())
	)

	var body = &activitypub.Create{
		Context:   activitypub.NewContext(),
		Id:        id,
		Type:      "Create",
		Actor:     h.aAddr,
		Published: time.Now().UTC(),
		To:        to,
		Cc:        cc,
		Object: struct {
			Id               string      `json:"id"`
			Type             string      `json:"type"`
			Title            string      `json:"title"`
			Summary          interface{} `json:"summary"`
			InReplyTo        interface{} `json:"inReplyTo"`
			Published        time.Time   `json:"published"`
			Url              string      `json:"url"`
			AttributedTo     string      `json:"attributedTo"`
			To               []string    `json:"to"`
			Cc               []string    `json:"cc"`
			Sensitive        bool        `json:"sensitive"`
			AtomUri          string      `json:"atomUri"`
			InReplyToAtomUri interface{} `json:"inReplyToAtomUri"`
			Conversation     string      `json:"conversation"`
			Content          string      `json:"content"`
			ContentMap       struct {
				En string `json:"en"`
			} `json:"contentMap"`
			Attachment []interface{} `json:"attachment"`
			Tag        []interface{} `json:"tag"`
			Replies    struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				First struct {
					Type   string        `json:"type"`
					Next   string        `json:"next"`
					PartOf string        `json:"partOf"`
					Items  []interface{} `json:"items"`
				} `json:"first"`
			} `json:"replies"`
		}{
			Id:               id,
			Type:             c.Type,
			Title:            c.Title,
			Summary:          c.Summary,
			InReplyTo:        nil,
			Published:        time.Time{},
			Url:              url,
			AttributedTo:     "",
			To:               to,
			Cc:               cc,
			Sensitive:        false,
			AtomUri:          "",
			InReplyToAtomUri: nil,
			Conversation:     "",
			Content:          c.Content,
			ContentMap: struct {
				En string `json:"en"`
			}{},
			Attachment: nil,
			Tag:        nil,
			Replies: struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				First struct {
					Type   string        `json:"type"`
					Next   string        `json:"next"`
					PartOf string        `json:"partOf"`
					Items  []interface{} `json:"items"`
				} `json:"first"`
			}{},
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

	// CREATE FOLLOW OUTBOX ...
	if err := outbox.NewOutboxes(h.actorId, id, h.inbox, activitypub.CreateType, string(marshal)).Create(); err != nil {
		return nil, err
	}

	return response(notok, ok)
}
