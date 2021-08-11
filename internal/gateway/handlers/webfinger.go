package handlers

import (
	"encoding/json"
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/disism/hvxahv/pkg/remote"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

func WebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")

	ok := activitypub.IsRemote(resource)
	if ok {
		// If you are not searching for the user of this instance, go to the remote request.
		body, err := remote.NewClient("GET", activitypub.NewWebFingerUrl(activitypub.GetHost(resource), resource)).Get()
		if err != nil {
			log.Println(err)
		}
		var r activitypub.WebFingerData
		_ = json.Unmarshal(body, &r)
		c.JSON(200, r)
		return
	}

	// Perform some filtering operations from the request to obtain the user name,
	// and then search for the user name to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.
	// Use this client to call the remote Accounts gRPC service,
	// and then pass the username to get the queried data.
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	accounts, err := cli.FindAccount(context.Background(), &pb.AccountByName{Username: activitypub.GetActorName(resource)})
	if err != nil {
		return
	}

	c.JSON(200, activitypub.NewWebFinger(accounts.Username))

}
