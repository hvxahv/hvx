package handlers

import (
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubHandler(c *gin.Context) {
	account, err := accounts.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		log.Println(err)
		return
	}
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")
	device, err := accounts.NewDeviceByHash(account.ID, middleware.GetDevicesID(c)).GetDeviceByHash()
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
