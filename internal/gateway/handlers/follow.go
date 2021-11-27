package handlers

import (
	"encoding/json"
	"fmt"
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

	f, addr := activity.NewFoAPData(name, uint(id))
	data, err := json.Marshal(f)
	if err != nil {
		return 
	}

	if err := activity.NewAPData(name, addr, data).Send(); err != nil {
		return 
	}
	c.JSON(200, gin.H{
		"code": 200,
		"message": "ok",
	})
}

func FollowAcceptHandler(c *gin.Context) {
	name := middleware.GetUserName(c)

	id := c.PostForm("id")
	actor := c.PostForm("actor")
	object := c.PostForm("object")

	aID, err := strconv.Atoi(actor)
	if err != nil {
		return
	}

	oID, err := strconv.Atoi(object)
	if err != nil {
		return
	}

	fa, addr := activity.NewFoAPAccept(name, id, uint(aID))
	data, err := json.Marshal(fa)
	if err != nil {
		return
	}

	fmt.Println(addr)

	if err := activity.NewAPData(name, addr, data).Send(); err != nil {
		return
	}

	// objectID (remote) -> actorID (user)
	if err := activity.NewFollows(uint(oID), uint(aID)).Create(); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"message": "ok",
	})
}
