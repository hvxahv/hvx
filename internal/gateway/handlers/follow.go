package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/activity"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"strconv"
)

func FollowReqHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	object := c.PostForm("object")
	id, err := strconv.Atoi(object)
	if err != nil {
		return 
	}

	f, inboxAddr := activity.NewFollowRequest(name, uint(id))
	data, err := json.Marshal(f)
	if err != nil {
		return 
	}

	if err := activity.NewActivityRequest(name, inboxAddr, data).Send(); err != nil {
		return 
	}
	c.JSON(200, gin.H{
		"code": 200,
		"message": "ok",
	})
}
