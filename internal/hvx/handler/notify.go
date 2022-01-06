package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"github.com/hvxahv/hvxahv/internal/notify"
	"log"
)

func NotifySubHandler(c *gin.Context) {
	a, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		log.Println(err)
		return
	}
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")
	device, err := device.NewDeviceByHash(a.ID, middleware.GetDevicesID(c)).GetDeviceByHash()
	if err != nil {

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
