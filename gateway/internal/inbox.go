package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"io"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	var body []byte
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
	}
	inbox, err := clientv1.New(c, microsvc.ActivityServiceName).Inbox(name, body)
	if err != nil {
		return
	}
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(200, inbox)

	//if err := proxy.NewProxy(c, "/u/"+c.Param("actor")+"/inbox", address.Public).Proxy(); err != nil {
	//	c.JSON(502, gin.H{
	//		"error": "502_BAD_GATEWAY",
	//	})
	//	return
	//}
}

func ChannelInboxHandler(c *gin.Context) {
	name := c.Param("channel")
	var body []byte
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
	}
	inbox, err := clientv1.New(c, microsvc.ActivityServiceName).Inbox(name, body)
	if err != nil {
		return
	}
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(200, inbox)

	//if err := proxy.NewProxy(c, "/c/"+c.Param("channel")+"/inbox", address.Public).Proxy(); err != nil {
	//	c.JSON(502, gin.H{
	//		"error": "502_BAD_GATEWAY",
	//	})
	//	return
	//}
}
