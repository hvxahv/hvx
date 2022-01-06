package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"github.com/hvxahv/hvxahv/internal/notify"
	"github.com/hvxahv/hvxahv/pkg/push"
	"github.com/hvxahv/hvxahv/pkg/security"
	"log"
	"strconv"
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
	//	log.Printf("failed to send message to account server: %v", err)
	//}

	id, mail, err := account.NewAuth(username, password).SignIn()
	if err != nil {
		fmt.Println(err)
		return
	}
	deviceID := uuid.New().String()
	token, err := security.GenToken(mail, username, password, deviceID)
	if err != nil {
		log.Println(err)
	}
	d := device.NewDevices(id, ua, deviceID)
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
	actor, err := account.NewActorByAccountUsername(name).GetActorByAccountUsername()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":       "200",
		"public_key": actor.PublicKey,
	})
}

func GetPrivateKeyHandlers(c *gin.Context) {
	account, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		log.Println(err)
		return
	}
	i, err := strconv.Atoi(c.Query("device_id"))
	if err != nil {
		log.Println(err)
		return
	}

	d, err := json.Marshal(push.NewData("Notify", fmt.Sprintf("%v: Signing in, requesting your private key.", i), "https://avatars.githubusercontent.com/u/94792300?s=200&v=4", "Normal"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := notify.NewPush(account.ID, uint(i), d).Push(); err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "The request is successful, please confirm login in the requested terminal.",
	})

}

func LogoutHandler(c *gin.Context) {
	acct, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := device.NewDeviceByHash(acct.ID, middleware.GetDevicesID(c)).Delete(); err != nil {
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
	acct, err := account.NewAccountsUsername(username).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	devices, err := device.NewDevicesByAccountID(acct.ID).GetDevicesByAccountID()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"devices": devices,
	})
}

func DeleteDevicesHandler(c *gin.Context) {
	acct, err := account.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := device.NewDeviceByHash(acct.ID, c.PostForm("device_hash")).Delete(); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
