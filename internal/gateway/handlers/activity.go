package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/activity"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))

	activity.Types(name, body)
}

func OutboxHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	t := c.PostForm("type")
	o := c.PostForm("object")

	fmt.Println(name, t, o)
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}

func GetInboxesHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	fmt.Println(name)
	acct, err := accounts.NewAccountsUsername(name).GetAccountByUsername()
	if err != nil {
		return
	}
	inboxes, err := activity.NewObjectID(acct.ActorID).GetInboxesByID()

	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": inboxes,
	})
}
