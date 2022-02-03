package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
)

// GetActorHandler Get the actor's handler. It will get the queried username from Param,
// then call the gRPC service by the username,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	client, err := account.NewActorClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountUsername{
		Username: c.Param("actor"),
	}

	actor, err := client.GetActorByAccountUsername(c, d)
	if err != nil {
		return
	}
	a := activitypub.NewActor(actor)
	c.JSON(200, a)
}

func SearchActorsHandler(c *gin.Context) {
	client, err := account.NewActorClient()
	if err != nil {
		return
	}
	d := &pb.NewActorPreferredUsername{
		PreferredUsername: c.Param("actor"),
	}

	actors, err := client.GetActorsByPreferredUsername(c, d)
	if err != nil {
		return
	}
	c.JSON(200, actors)
}

func EditActorHandler(c *gin.Context) {
	name := c.PostForm("name")
	summary := c.PostForm("summary")
	avatar := c.PostForm("avatar")

	client, err := account.NewActorClient()
	if err != nil {
		return
	}
	d := &pb.NewEditActor{
		AccountUsername: middleware.GetUsername(c),
		Avatar:          avatar,
		Name:            name,
		Summary:         summary,
	}
	reply, err := client.EditActor(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": reply.Reply,
	})
}
