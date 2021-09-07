package handlers

import (
	"fmt"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/internal/messages"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, _ := ioutil.ReadAll(c.Request.Body)

	messages.InboxEventHandler(name, body)
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
