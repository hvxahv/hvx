package activity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func NewDelete(articleID uint) *activitypub.Delete {
	actor := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), "hvturingga")
	id := fmt.Sprintf("%s/article/%s", actor, strconv.Itoa(int(articleID)))
	to := []string{"https://mas.to/users/hvturingga/"}

	return &activitypub.Delete{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      fmt.Sprintf("%s#delete", id),
		Type:    "Delete",
		Actor:   actor,
		To:      to,
		Object: struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			AtomUri string `json:"atomUri"`
		}{
			Id:      id,
			Type:    "Tombstone",
			AtomUri: id,
		},
	}
}

func NewArticle(articleID uint, content string) *activitypub.Create {
	actor := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), "hvturingga")
	id := fmt.Sprintf("%s/article/%s", actor, strconv.Itoa(int(articleID)))
	to := []string{"https://mas.to/users/hvturingga/"}
	return &activitypub.Create{
		Context:   "https://www.w3.org/ns/activitystreams",
		Id:        fmt.Sprintf("%s/activity", id),
		Type:      "Create",
		Actor:     actor,
		Published: time.Time{},
		To:        to,
		Cc:        nil,
		Object: struct {
			Id               string      `json:"id"`
			Type             string      `json:"type"`
			Summary          interface{} `json:"summary"`
			InReplyTo        string      `json:"inReplyTo"`
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
			Attachment       []struct {
				Type      string      `json:"type"`
				MediaType string      `json:"mediaType"`
				Url       string      `json:"url"`
				Name      interface{} `json:"name"`
				Blurhash  string      `json:"blurhash"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
			} `json:"attachment"`
			Tag     []interface{} `json:"tag"`
			Replies struct {
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
			Type:             "Note",
			Summary:          nil,
			InReplyTo:        "",
			Published:        time.Time{},
			Url:              id,
			AttributedTo:     "",
			To:               to,
			Cc:               nil,
			Sensitive:        false,
			AtomUri:          "",
			InReplyToAtomUri: nil,
			Conversation:     "",
			Content:          content,
			Attachment:       nil,
			Tag:              nil,
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
	}
}

func NewFollowRequest(actor string, objectID uint) (*activitypub.Follow, string) {
	o, err := accounts.NewActorID(objectID).GetByID()
	if err != nil {
		return nil, ""
	}

	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/%s", viper.GetString("localhost"), uuid.New().String())
		a   = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor)
	)

	return &activitypub.Follow{
		Context: ctx,
		Id:      id,
		Type:    "Follow",
		Actor:   a,
		Object:  o.Url,
	}, o.Inbox
}
