package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/internal/messages"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, _ := ioutil.ReadAll(c.Request.Body)

	inbox := messages.Inbox{}
	err := json.Unmarshal(body, &inbox)
	if err != nil {
		fmt.Println(err)
	}

	inbox.Inbox(name)
}

func OutboxHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	t := c.PostForm("type")
	o := c.PostForm("object")

	fmt.Println(name, t, o)
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
