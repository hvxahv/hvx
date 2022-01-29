package handler

import (
	"github.com/gin-gonic/gin"
)

func NotifySubHandler(c *gin.Context) {
	//a, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//endpoint := c.PostForm("endpoint")
	//p256dh := c.PostForm("p256dh")
	//auth := c.PostForm("auth")
	//d, err := device.NewDeviceByHash(a.ID, middleware.GetDevicesID(c)).GetDeviceByHash()
	//if err != nil {
	//
	//}
	//if err := notify.NewNotifies(d.ID, endpoint, p256dh, auth).Create(); err != nil {
	//	log.Panicln(err)
	//	return
	//}
	//c.JSON(200, gin.H{
	//	"code":    "200",
	//	"message": "ok",
	//})
}
