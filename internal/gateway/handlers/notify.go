package handlers

import (
	"github.com/hvxahv/hvxahv/internal/accounts"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubHandler(c *gin.Context) {
	//  deviceID uint, endpoint, p256dh, auth, public_key, private_key string
	// GET DEVICE ID BY ACCOUNT ID.
	deviceID := c.PostForm("device_id")
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")

	device, err := accounts.NewDevicesID(deviceID).GetDevicesByDeviceID()
	if err != nil {
		log.Println(err)
		return
	}

	if err := notify.NewNotifies(device.ID, endpoint, p256dh, auth).Create(); err != nil {
		log.Panicln(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok",
	})
}
