package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/pkg/security"
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
		return
	}
	deviceID := uuid.New().String()
	token, err := security.GenToken(mail, username, password, deviceID)
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

func LogoutHandler(c *gin.Context) {
	username := middleware.GetUsername(c)
	devices := middleware.GetDevicesID(c)

	acct, err := accounts.NewAccountsUsername(username).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := accounts.NewDevicesByDeviceID(acct.ID, devices).DeleteByDeviceID(); err != nil {
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
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.PostForm("id"))
	fmt.Println(id)
	acct, err := accounts.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := accounts.NewDevicesByID(uint(id), acct.ID).DeleteByID(); err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
