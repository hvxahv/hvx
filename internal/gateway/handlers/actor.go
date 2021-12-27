package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/activity"
	"github.com/hvxahv/hvxahv/internal/channels"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/spf13/viper"
	"io/ioutil"
	"time"
)

// GetActorHandler Get the actor's handler. It will get the queried username from Param,
// then call the gRPC service by the username,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	name := c.Param("actor")

	// Use this client to call the remote Accounts gRPC service,
	// and then pass the username to get the queried data.
	//cli, conn, err := client.Accounts()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer conn.Close()
	//acct, err := cli.FindActorByAccountsUsername(context.Background(), &pb.AccountUsername{Username: name})
	//if err != nil {
	//	return
	//}
	//
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"status":  "600",
	//		"message": "NO QUERY TO THE ACCOUNT.",
	//	})
	//}

	acct, err := accounts.NewActorsPreferredUsername(name).GetActorByAccountUsername()
	if err != nil {
		return
	}

	a := NewActor(acct)
	c.JSON(200, a)
}

func GetChannelHandler(c *gin.Context) {
	name := c.Param("actor")
	ch, err := channels.NewChannelsByLink(name).GetActorDataByLink()
	if err != nil {
		return
	}
	a := NewChannelActor(ch)
	c.JSON(200, a)
}

func ChannelInboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	activity.ChannelTypes(name, body)
}

func NewContext() []interface{} {
	arr := make([]interface{}, 0)
	ctx := []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1alpha1"}
	for _, i := range ctx {
		arr = append(arr, i)
	}
	return arr
}

func NewChannelActor(a *accounts.Actors) *activitypub.Actor {
	var (
		addr = viper.GetString("localhost")

		id  = fmt.Sprintf("https://%s/c/%s", addr, a.PreferredUsername)
		kid = fmt.Sprintf("%s#main-key", id)
		box = fmt.Sprintf("https://%s/c/%s/", addr, a.PreferredUsername)
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
		PreferredUsername:         a.PreferredUsername,
		Name:                      a.Name,
		Summary:                   a.Summary,
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

// NewActor Return standard ActivityPub protocol user data.
func NewActor(a *accounts.Actors) *activitypub.Actor {
	var (
		addr = viper.GetString("localhost")

		id  = fmt.Sprintf("https://%s/u/%s", addr, a.PreferredUsername)
		kid = fmt.Sprintf("%s#main-key", id)
		box = fmt.Sprintf("https://%s/u/%s/", addr, a.PreferredUsername)
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
		PreferredUsername:         a.PreferredUsername,
		Name:                      a.Name,
		Summary:                   a.Summary,
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
