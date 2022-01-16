package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/channel"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"log"
)

func WebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")
	ok := activitypub.IsRemote(resource)
	if ok {
		actor := activitypub.GetWebFinger(resource)
		c.JSON(200, actor)
		return
	}

	// Perform some filtering operations from the request to obtain the user name,
	// and then search for the user name to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.
	// Use this client to call the remote Accounts gRPC service,
	// and then pass the username to get the queried data.

	acct, err := account.NewAccountsUsername(activitypub.GetActorName(resource)).GetAccountByUsername()
	log.Println(acct)
	if err != nil {
		ch, err := channel.NewChannelsByLink(activitypub.GetActorName(resource)).GetActorDataByLink()
		if err != nil {
			return
		}
		c.JSON(200, activitypub.NewWebFinger(ch.PreferredUsername, true))
		return
	}

	c.JSON(200, activitypub.NewWebFinger(acct.Username, false))

}
