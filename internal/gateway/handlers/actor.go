package handlers

import (
	"github.com/disism/hvxahv/pkg/activitypub"
	"github.com/gin-gonic/gin"
)

// GetActorHandler Get the actor's handler. It will get the queried user name from Param,
// then call the gRPC service by the user name,
// and return the JsonLD of the standard activitypub protocol.
func GetActorHandler(c *gin.Context) {
	name := c.Param("actor")
	acct := GetAccounts(name)
	a := activitypub.NewActor(acct)
	c.JSON(200, a)
}
