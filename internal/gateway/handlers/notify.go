package handlers

import (
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubHandler(c *gin.Context) {
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
