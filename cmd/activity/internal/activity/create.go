package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cmd/activity/internal/delivery"
	"time"
)

type ArticleCreateResponse struct {
	Status  string
	Address string
}

func NewArticleCreateResponse(status string, address string) *ArticleCreateResponse {
	return &ArticleCreateResponse{Status: status, Address: address}
}

func (h *Handler) Create(address string, in *pb.ArticleCreateActivityRequest) (*ArticleCreateResponse, error) {

	// TODO - CREATE ARTICLE...
	var (
		id  = fmt.Sprintf("%s/articles/%s/activity#create", h.Actor.Address, uuid.NewString())
		url = fmt.Sprintf("%s/articles/%s/activity#create", h.Actor.Address, uuid.NewString())
	)
	if len(in.Article.GetTo()) < 1 {
		in.Article.To = append(in.Article.To, fmt.Sprintf("%s/followers", h.Actor.Address))
	}
	if len(in.Article.GetAudience()) > 0 {
		in.Article.Cc = append(in.Article.To, fmt.Sprintf("%s/followers", h.Actor.Address))
	}
	var body = &activitypub.Create{
		Context:   activitypub.NewContext(),
		Id:        id,
		Type:      activitypub.CreateType,
		Actor:     h.Actor.Address,
		Published: time.Now().UTC(),
		To:        in.Article.GetTo(),
		Cc:        in.Article.GetCc(),
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
			Type:             activitypub.NoteType,
			Title:            in.Article.GetSummary(),
			Summary:          in.Article.Summary,
			InReplyTo:        nil,
			Published:        time.Time{},
			Url:              url,
			AttributedTo:     "",
			To:               in.Article.To,
			Cc:               in.Article.Cc,
			Sensitive:        false,
			AtomUri:          "",
			InReplyToAtomUri: nil,
			Conversation:     "",
			Content:          in.Article.GetArticle(),
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
	do, err := delivery.New(h.Actor.PublicKeyId, h.Actor.PrivateKey, marshal).Do(h.Object.Inbox)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != 202 {
		return NewArticleCreateResponse("failures", address), nil
	}

	return NewArticleCreateResponse("successes", address), nil
}
