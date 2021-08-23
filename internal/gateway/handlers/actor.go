package handlers

import (
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"strings"
)

// GetActorHandler Get the actor's handler. It will get the queried user name from Param,
// then call the gRPC service by the user name,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	name := strings.Trim(c.Param("actor"), ".json")

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
	a := activitypub.NewActor(acct)
	c.JSON(200, a)
}
