package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/channels"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
)

func WebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")

	//cli, conn, err := client.Accounts()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer conn.Close()

	// Determine whether it is a user of this instance,
	// if you do not search for a user of this instance, go to the remote request.
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

	//accounts, err := cli.GetAccountsByUsername(context.Background(), &pb.AccountUsername{Username: activitypub.GetActorName(resource)})
	//if err != nil {
	//	ch, err := channels.NewChannelsByLink(activitypub.GetActorName(resource)).GetActorDataByLink()
	//	if err != nil {
	//		return
	//	}
	//	c.JSON(200, activitypub.NewWebFinger(ch.PreferredUsername, true))
	//	return
	//}
	acct, err := accounts.NewAccountsUsername(activitypub.GetActorName(resource)).GetAccountByUsername()
	if err != nil {
		ch, err := channels.NewChannelsByLink(activitypub.GetActorName(resource)).GetActorDataByLink()
		if err != nil {
			return
		}
		c.JSON(200, activitypub.NewWebFinger(ch.PreferredUsername, true))
		return
	}

	c.JSON(200, activitypub.NewWebFinger(acct.Username, false))

}
