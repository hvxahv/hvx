package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
)

func WebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")
	ok := activitypub.IsRemote(resource)
	if ok {
		actor := activitypub.GetWebFinger(resource)
		c.JSON(200, actor)
		return
	}

	// Perform some filtering operations from the request to obtain the username,
	// and then search for the username to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.
	// Use this client to call the remote Accounts gRPC service,
	// and then pass the username to get the queried data.

	client, err := account.NewClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountUsername{Username: activitypub.GetActorName(resource)}
	a, err := client.GetAccountByUsername(c, d)
	if err != nil {
		return
	}

	//if err != nil {
	//	ch, err := channel.NewChannelsByLink(activitypub.GetActorName(resource)).GetActorDataByLink()
	//	if err != nil {
	//		return
	//	}
	//	c.JSON(200, activitypub.NewWebFinger(ch.PreferredUsername, true))
	//	return
	//}

	c.JSON(200, activitypub.NewWebFinger(a.Username, false))

}
