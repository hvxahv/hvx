package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/gateway/address"
	"github.com/hvxahv/hvx/gateway/proxy"
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
	client, err := clientv1.New(c, microsvc.NewGRPCAddress("activity").Get())
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
		return
	}
	defer client.Close()

	inbox, err := activity.NewInboxClient(client.Conn).Inbox(c, &activity.InboxRequest{
		Name: name,
		Data: body,
	})
	if err != nil {
		c.JSON(502, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(200, inbox)

	if err := proxy.NewProxy(c, "/u/"+c.Param("actor")+"/inbox", address.Public).Proxy(); err != nil {
		c.JSON(502, gin.H{
			"error": "502_BAD_GATEWAY",
		})
		return
	}
}
