package handlers

import (
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubHandler(c *gin.Context) {
	// deviceID uint, endpoint, p256dh, auth, public_key, private_key string

	// Use web push API to implement web message push interface.
	// https://datatracker.ietf.org/doc/html/rfc8292
	// https://developer.mozilla.org/en-US/docs/Web/API/Push_API
	// https://datatracker.ietf.org/doc/html/rfc6108

	deviceID := middleware.GetDevicesID(c)
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")

	if err := notify.NewNotifies(deviceID, endpoint, p256dh, auth).Create(); err != nil {
		log.Panicln(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok",
	})
}
