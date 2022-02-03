package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubscriptionHandler(c *gin.Context) {
	id := c.PostForm("device_id")
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")

	fmt.Println(id, endpoint, p256dh, auth)
	
	client, err := notify.NewNotifyClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	d := &v1alpha1.NewNotifySubscription{
		DeviceId: id,
		Endpoint: endpoint,
		P256Dh:   p256dh,
		Auth:     auth,
	}

	reply, err := client.Subscription(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  reply.Code,
		"reply": reply.Reply,
	})
}
