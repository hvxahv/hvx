package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/activitypub"
	"hvxahv/pkg/microservices/client"
	"log"
)

// WebFinger and WebFingerLinks form the JSON-LD protocol of the standard Activitypub
type WebFinger struct {
	Subject string           `json:"subject"`
	Links   []WebFingerLinks `json:"links"`
}

// WebFingerLinks is used by WebFinger
type WebFingerLinks struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

func WebFingerHandler(c *gin.Context) {
	// Perform some filtering operations from the request to obtain the user name,
	// and then search for the user name to find whether the user exists in the database.
	// Currently only tested mastodon has not supported other ActivityPub implementations.
	res := c.Query("resource")
	name, err := activitypub.IsActorExists(res)
	if err != nil {
		log.Println(err)
		return
	}

	// Use this client to call the remote Accounts gRPC service,
	// and then pass the user name to get the queried data.
	cli, conn,  err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	accounts, err := cli.QueryAccounts(context.Background(), &pb.AccountsName{Username: name})
	if err != nil {
		return 
	}


	c.JSON(200, activitypub.NewWebFinger(accounts.Username))
}

