package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
)

func GetWebFingerHandler(c *gin.Context) {
	resource := c.Query("resource")
	ok := activitypub.IsRemote(resource)
	if ok {
		actor := activitypub.GetRemoteWebFinger(resource)
		c.JSON(200, actor)
		return
	}

	name := activitypub.GetActorName(resource)
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	
	d := &pb.NewAccountUsername{Username: name}

	e, err := client.IsExist(c, d)
	if err != nil {
		return
	}

	//if err != nil {
	//	ch, err := channel.NewChannelsByLink(activitypub.GetActorName(resource)).GetActorDataByLink()
	//	if err != nil {
	//		return
	//	}
	//	c.JSON(200, activitypub.NewWebFinger(ch.PreferredUsername, true))
	//	return
	//}

	if e.IsExist {
		c.JSON(200, activitypub.NewWebFinger(name, false))
		return
	}
}
