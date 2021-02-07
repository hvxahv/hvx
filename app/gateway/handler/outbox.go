package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/cli/social"
	"hvxahv/pkg/models"
	"hvxahv/pkg/utils"
)

func FollowersAcceptHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	actor := c.PostForm("actor")
	id := c.PostForm("id")

	data := models.NewAccept(actor, name, id)
	social.OutboxAcceptClient(data)

}


