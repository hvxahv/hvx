package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
)

func CreateAccountHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	mail := c.PostForm("mail")

	// https://datatracker.ietf.org/doc/html/rfc5208
	publicKey := c.PostForm("public_key")

	d := &pb.NewAccountCreate{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	}

	cli, err := account.NewAccountClient()
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	create, err := cli.Create(c, d)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  create.Code,
		"reply": create.Reply,
	})

}

func SignInHandler(c *gin.Context) {
	cli, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountVerify{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Ua:       c.GetHeader("User-Agent"),
	}

	verify, err := cli.Verify(c, d)
	if err != nil {
		c.JSON(401, gin.H{
			"code":  "401",
			"reply": "Unauthorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":       "200",
		"token":      verify.Token,
		"mail":       verify.Mail,
		"device_id":  verify.DeviceId,
		"public_key": verify.PublicKey,
	})
}

func LogoutHandler(c *gin.Context) {
	client, err := device.NewDeviceClient()
	if err != nil {
		return
	}
	d := &v1alpha1.NewDeviceHash{
		Hash: middleware.GetDeviceHash(c),
	}
	reply, err := client.DeleteByDeviceHash(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  reply.Code,
		"reply": reply.Reply,
	})
}

func GetPublicKeyHandlers(c *gin.Context) {
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountUsername{
		Username: middleware.GetUsername(c),
	}
	reply, err := client.GetPublicKeyByAccountUsername(c, d)
	c.JSON(200, gin.H{
		"code":       reply.Code,
		"public_key": reply.PublicKey,
	})
}
