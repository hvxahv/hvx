package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/activitypub"
	"hvxahv/pkg/microservices/client"
	"log"
)

// GetActorHandler Get the actor's handler. It will get the queried user name from Param,
// then call the gRPC service by the user name,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	name := c.Param("actor")

	// Use the client to call the Accounts service to create users.
	// Pass in the username and search for the user, if found, the accounts data will be returned.
	cli, conn,  err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	accounts, err := cli.QueryAccounts(context.Background(), &pb.AccountsName{Username: name})
	if err != nil {
		return
	}

	a := activitypub.NewActor(accounts)
	c.JSON(200, a)
}
