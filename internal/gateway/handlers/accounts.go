package handlers

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

// NewAccountsHandler ...
func NewAccountsHandler(c *gin.Context) {
	// Username used to log in.
	username := c.PostForm("username")
	// Password for login.
	password := c.PostForm("password")

	mail := c.PostForm("mail")

	// Use the client to call the Accounts service to create users.
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewAccount(context.Background(), &pb.NewAccountData{
		Username: username,
		Password: password,
		Mail:     mail,
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}
	fmt.Println(r)
	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})

}

// LoginHandler ...
func LoginHandler(c *gin.Context) {
	mail := c.PostForm("mail")
	// Password for login.
	password := c.PostForm("password")

	// Use this client to remotely call the login method.
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	r, err := cli.LoginAccount(context.Background(), &pb.LoginData{
		Mail: mail,
		Password: password,
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}
	c.JSON(200, gin.H{
		"code":     "200",
		"username": r.Username,
		"token":    r.Token,
	})
}
//
//// UploadAvatar Interface for users to upload avatars.
//func UploadAvatar(c *gin.Context) {
//	file, err := c.FormFile("file")
//	if err != nil {
//		fmt.Println("File read failed！" + err.Error())
//		c.JSON(500, gin.H{
//			"status":  500,
//			"message": "File read failed！",
//		})
//		return
//	}
//
//	bucket := "accounts"
//	m := oos.NewMin(file, "avatar", bucket, "")
//	url, err := m.FileUploader()
//	if err != nil {
//		log.Println(err)
//		c.JSON(500, gin.H{
//			"status":  500,
//			"message": "Server receiving file error!",
//		})
//		return
//	}
//	var link string
//	if viper.GetBool("oos.minio.useSSL") {
//		link = fmt.Sprintf("https://%s/%s/%s", viper.GetString("oos.minio.addr"), bucket, url)
//	} else {
//		link = fmt.Sprintf("http://%s/%s/%s", viper.GetString("oos.minio.addr"), bucket, url)
//	}
//	c.JSON(200, gin.H{
//		"status":  200,
//		"message": "Avatar uploaded successfully.",
//		"url":     link,
//	})
//}
//
//func GetAccountsHandler(c *gin.Context) {
//	//name := middleware.GetUserName(c)
//	//a := GetAccounts(name)
//	//log.Println(name)
//	//log.Println(a)
//
//	c.JSON(200, gin.H{
//		"code":     200,
//		"messages": "ok",
//	})
//}
//
//// GetAccountsByName Incoming username is used to query account.
//func GetAccountsByName(name string) (*pb.AccountsData, error) {
//	// Use the client to call the Accounts service to create users.
//	// Pass in the username and search for the user, if found, the accounts data will be returned.
//	client, conn, err := client.Accounts()
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//	accounts, err := client.QueryAccounts(context.Background(), &pb.AccountsData{Username: name})
//	if err != nil {
//		return nil, err
//	}
//
//	return accounts, nil
//}
