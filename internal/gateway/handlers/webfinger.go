package handlers

import (
	"encoding/json"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
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
		wf, err := remote.NewClient("GET", activitypub.NewWebFingerUrl(activitypub.GetHost(resource), resource)).Get()
		if err != nil {
			log.Println(err)
		}
		var wfd activitypub.WebFingerData
		_ = json.Unmarshal(wf, &wfd)


		var ar activitypub.Actor
		actor, err := remote.NewClient("GET", wfd.Links[0].Href).Get()
		if err != nil {
			log.Println(err)
		}
		_ = json.Unmarshal(actor, &ar)

		c.JSON(200, ar)
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

	accounts, err := cli.Find(context.Background(), &pb.NewAccountByName{Username: activitypub.GetActorName(resource)})
	if err != nil {
		return
	}

	c.JSON(200, activitypub.NewWebFinger(accounts.Username))

}
