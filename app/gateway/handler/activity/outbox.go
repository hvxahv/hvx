package activity

import (
	"github.com/gin-gonic/gin"
	social2 "hvxahv/pkg/api-client/social"
	"hvxahv/pkg/models"
	"hvxahv/pkg/response"
	"hvxahv/pkg/utils"
	"log"
)

// FollowersAcceptHandler 同意关注的 Handler
func FollowerAcceptHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	actor := c.PostForm("actor")
	id := c.PostForm("id")

	data := models.NewAccept(actor, name, id)
	r, err := social2.OutboxAcceptClient(data)
	if err != nil {
		log.Println(err)
	}

	response.SendActivityResponse(c, r.Reply)
}

// FollowHandler ... 请求关注的 Handler
func FollowHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	actor := c.PostForm("actor")

	r, err := social2.OutboxFollowClient(name, actor)
	if err != nil {
		log.Println(err)
	}

	response.SendActivityResponse(c, r.Reply)
}

