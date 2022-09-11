package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cfg"
	"testing"
)

func init() {
	cfg.DefaultConfig()
}

func TestDelivery_Do(t *testing.T) {
	const (
		actorAddress  = "https://halfmemories.com/u/hvturingga"
		inbox         = "https://mastodon.disism.com/users/hvturingga/inbox"
		object        = "https://mastodon.disism.com/users/hvturingga"
		pemPrivateKey = ``
	)
	body := &activitypub.Follow{
		Context: "https://www.w3.org/ns/activitystreams",
		Id:      fmt.Sprintf("%s/%s", actorAddress, uuid.NewString()),
		Type:    "Follow",
		Actor:   actorAddress,
		Object:  object,
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(marshal)
}
