package handlers

import (
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/notify"
)

func NotifySubRequestHandler(c *gin.Context) {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"PrivateKey": privateKey,
		"PublicKey":  publicKey,
	})
}

func NotiftSubHandler(c *gin.Context) {
	//  deviceID uint, endpoint, p256dh, auth, public_key, private_key string
	// GET DEVICE ID BY ACCOUNT ID.
	account_id := c.PostForm("id")
	endpoint := c.PostForm("endpoint")
	p256dh := c.PostForm("p256dh")
	auth := c.PostForm("auth")
	public_key := c.PostForm("public_key")
	private_key := c.PostForm("private_key")

	fmt.Println(account_id)

	if err := notify.NewNotifies(1123123124, endpoint, p256dh, auth, public_key, private_key).Create(); err != nil {
		log.Panicln(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok",
	})
}
