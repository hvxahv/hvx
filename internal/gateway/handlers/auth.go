package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/pkg/security"
	"log"
)

func SignInHandler(c *gin.Context) {
	ua := c.GetHeader("User-Agent")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Use this client to remotely call the login method.
	//cli, conn, err := client.Accounts()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer conn.Close()
	//r, err := cli.SignIn(context.Background(), &pb.AuthData{
	//	Username: username,
	//	Password: password,
	//})
	//if err != nil {
	//	log.Printf("failed to send message to accounts server: %v", err)
	//}

	id, mail, err := accounts.NewAuth(username, password).SignIn()
	if err != nil {
		fmt.Println(err)
		return
	}
	deviceID := uuid.New().String()
	token, err := security.GenToken(mail, username, password, deviceID)
	if err != nil {
		log.Println(err)
	}
	d := accounts.NewDevices(id, ua, deviceID)
	if err := d.Create(); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":      "200",
		"token":     token,
		"mail":      mail,
		"deviceID":  deviceID,
		"publicKey": d.PublicKey,
	})
}

func GetPublicKeyHandlers(c *gin.Context) {
	name := middleware.GetUsername(c)
	actor, err := accounts.NewActorByAccountUsername(name).GetActorByAccountUsername()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":       "200",
		"public_key": actor.PublicKey,
	})
}

func LogoutHandler(c *gin.Context) {
	username := middleware.GetUsername(c)
	devices := middleware.GetDevicesID(c)

	acct, err := accounts.NewAccountsUsername(username).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := accounts.NewDevicesByAccountIDAndDeviceID(acct.ID, devices).DeleteByDeviceID(); err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "logout",
	})
}

func GetDevicesHandler(c *gin.Context) {
	username := middleware.GetUsername(c)
	acct, err := accounts.NewAccountsUsername(username).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	devices, err := accounts.NewDevicesByAccountID(acct.ID).GetDevices()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"devices": devices,
	})
}

func DeleteDevicesHandler(c *gin.Context) {
	deviceID := c.PostForm("device_id")
	acct, err := accounts.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := accounts.NewDevicesByAccountIDAndDeviceID(acct.ID, deviceID).DeleteByDeviceID(); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
