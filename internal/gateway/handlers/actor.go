package handlers

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"time"
)

// GetActorHandler Get the actor's handler. It will get the queried user name from Param,
// then call the gRPC service by the user name,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	name := c.Param("actor")

	// Use this client to call the remote Accounts gRPC service,
	// and then pass the user name to get the queried data.
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	acct, err := cli.Find(context.Background(), &pb.NewAccountByName{Username: name})
	if err != nil {
		return
	}

	if err != nil {
		c.JSON(200, gin.H{
			"status":  "600",
			"message": "NO QUERY TO THE ACCOUNT.",
		})
	}

	a := NewActor(acct)
	c.JSON(200, a)
}


func NewContext() []string {
	ctx := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1alpha1"}
	return ctx
}

// NewActor Return standard ActivityPub protocol user data.
func NewActor(a *pb.AccountData) *activitypub.Actor {
	var (
		addr = viper.GetString("localhost")

		id = fmt.Sprintf("https://%s/u/%s", addr, a.Username)
		kid = fmt.Sprintf("%s#main-key", id)
		box = fmt.Sprintf("https://%s/u/%s/", addr, a.Username)
	)

	actor := &activitypub.Actor{
		Context:                   NewContext(),
		Id:                        id,
		Type:                      "Person",
		Following:                 "",
		Followers:                 "",
		Inbox:                     box + "inbox",
		Outbox:                    box + "outbox",
		Featured:                  "",
		FeaturedTags:              "",
		PreferredUsername:         a.Username,
		Name:                      a.Name,
		Summary:                   a.Bio,
		Url:                       "",
		ManuallyApprovesFollowers: false,
		Discoverable:              false,
		Published:                 time.Time{},
		Devices:                   "",
		PublicKey: struct {
			Id           string `json:"id"`
			Owner        string `json:"owner"`
			PublicKeyPem string `json:"publicKeyPem"`
		}{
			Id:           kid,
			Owner:        id,
			PublicKeyPem: a.PublicKey,
		},
		Tag:        nil,
		Attachment: nil,
		Endpoints: struct {
			SharedInbox string `json:"sharedInbox"`
		}{},
		Icon: struct {
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
			Url       string `json:"url"`
		}{
			Type:      "Image",
			MediaType: "image/jpg",
			Url:       a.Avatar,
		},
	}
	return actor
}
