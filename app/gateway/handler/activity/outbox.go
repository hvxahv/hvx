package activity

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/middleware"
	"hvxahv/internal/activity"
	social2 "hvxahv/internal/client/social"
	"hvxahv/internal/outbox"
	"log"
)

// FollowersAcceptHandler 同意关注的 Handler
func FollowerAcceptHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	actor := c.PostForm("actor")
	id := c.PostForm("id")

	data := inbox.NewAccept(actor, name, id)
	r, err := social2.OutboxAcceptClient(data)
	if err != nil {
		log.Println(err)
	}

	activity.SendActivityResponse(c, r.Reply)
}

// FollowHandler ... 请求关注的 Handler
func FollowHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	actor := c.PostForm("actor")

	r, err := social2.OutboxFollowClient(name, actor)
	if err != nil {
		log.Println(err)
	}

	activity.SendActivityResponse(c, r.Reply)
}

