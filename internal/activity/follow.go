package activity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/spf13/viper"
	"log"
)

// NewFollowAccept
// name: LOCAL ACTOR NAME,
// actor: REMOTE ACTOR LINK,
// oid: CONTEXT ID,
// object: LOCAL ACTOR LINK.
func NewFollowAccept(name, object, activityID string, remoteActorID, localActorID uint) *activitypub.Accept {
	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/u/%s#accepts/follows/%s", viper.GetString("localhost"), name, uuid.New().String())
	)

	nf := accounts.NewFollows(remoteActorID, localActorID)
	if err := nf.New(); err != nil {
		log.Println(err)
	}

	return &activitypub.Accept{
		Context: ctx,
		Id:      id,
		Type:    "Accept",
		Actor:   object,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     activityID,
			Type:   "Follow",
			Actor:  "",
			Object: object,
		},
	}
}


func FollowAccept(id uint, name string) {
	db := cockroach.GetDB()
	var ibx Inboxes
	if err := db.Debug().Table("inboxes").Where("id = ?", id).First(&ibx).Error; err != nil {
		log.Println(err)
	}

	actor := "https://mas.to/users/hvturingga"

	object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)


	na := NewFollowAccept(name, object, ibx.ActivityID,  ibx.ActorID, ibx.LocalActorID)

	data, err := json.Marshal(na)
	if err != nil {
		log.Println(err)
		return
	}

	nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
	nar.Accept()
}