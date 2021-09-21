package activity

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

type Notes struct {

}

func NewCreateNote() *activitypub.Create  {
	localhost := viper.GetString("localhost")
	actor := fmt.Sprintf("https://%s/u/hvturingga", localhost)
	id := fmt.Sprintf("https://%s/%s", localhost, uuid.New().String())
	//to := []string{"https://mas.to/users/hvturingga"}

	return &activitypub.Create{
		Context:   nil,
		Id:        id,
		Type:      "Create",
		Actor:     actor,
		Published: time.Time{},
		To:        nil,
		Cc:        nil,
		Object: struct {
			Id               string        `json:"id"`
			Type             string        `json:"type"`
			Summary          interface{}   `json:"summary"`
			InReplyTo        interface{}   `json:"inReplyTo"`
			Published        time.Time     `json:"published"`
			Url              string        `json:"url"`
			AttributedTo     string        `json:"attributedTo"`
			To               []string      `json:"to"`
			Cc               []interface{} `json:"cc"`
			Sensitive        bool          `json:"sensitive"`
			AtomUri          string        `json:"atomUri"`
			InReplyToAtomUri interface{}   `json:"inReplyToAtomUri"`
			Conversation     string        `json:"conversation"`
			Content          string        `json:"content"`
			Attachment       []interface{} `json:"attachment"`
			Tag              []interface{} `json:"tag"`
			Replies          struct {
				Id    string `json:"id"`
				Type  string `json:"type"`
				First struct {
					Type   string        `json:"type"`
					Next   string        `json:"next"`
					PartOf string        `json:"partOf"`
					Items  []interface{} `json:"items"`
				} `json:"first"`
			} `json:"replies"`
		}{},
	}
}
