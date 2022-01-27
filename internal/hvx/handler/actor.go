package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
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

//func GetChannelHandler(c *gin.Context) {
//	name := c.Param("actor")
//	ch, err := channel.NewChannelsByLink(name).GetActorDataByLink()
//	if err != nil {
//		return
//	}
//	a := NewChannelActor(ch)
//	c.JSON(200, a)
//}
//
//func ChannelInboxHandler(c *gin.Context) {
//	//name := c.Param("actor")
//	body, err := ioutil.ReadAll(c.Request.Body)
//	if err != nil {
//		return
//	}
//	fmt.Println(string(body))
//	//activity.ChannelTypes(name, body)
//}

func SearchActorsHandler(c *gin.Context) {

}
