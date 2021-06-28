package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/microservices/client"
	"hvxahv/pkg/oos"
	"log"
	"strconv"
)

// NewAccountsHandler ...
func NewAccountsHandler(c *gin.Context) {
	// Username used to log in.
	username := c.PostForm("username")
	// Password for login.
	password := c.PostForm("password")
	// Account avatar.
	avatar := c.PostForm("avatar")
	// User's name, displayed name.
	name := c.PostForm("name")
	// User's email, used to retrieve password.
	email := c.PostForm("email")
	// Choose whether the account is a private account.
	p := c.PostForm("private")
	private, err := strconv.Atoi(p)
	if err != nil {
		log.Println(err)
	}

	// Use the client to call the Accounts service to create users.
	cli, conn,  err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewAccounts(context.Background(), &pb.NewAccountsData{
		Username: username,
		Password: password,
		Avatar:   avatar,
		Name:     name,
		Email:    email,
		Private:  int32(private),
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
	// Username used to log in.
	username := c.PostForm("username")
	// Password for login.
	password := c.PostForm("password")

	// Use this client to remotely call the login method.
	cli, conn,  err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	r, err := cli.LoginAccounts(context.Background(), &pb.AccountsLogin{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}
	c.JSON(200, gin.H{
		"code": "200",
		"username": r.Username,
		"token": r.Token,
	})
}


// UploadAvatar Interface for users to upload avatars.
func UploadAvatar(c *gin.Context)  {
	file, err := c.FormFile("file")
	if err!=nil {
		fmt.Println("File read failed！"+err.Error())
		c.JSON(500, gin.H{
			"status": 500,
			"message":"File read failed！",
		})
		return
	}

	bucket := "accounts"
	m := oos.NewMin(file, "avatar", bucket, "")
	url, err := m.FileUploader()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status": 500,
			"message":"Server receiving file error!",
		})
		return
	}
	var link string
	if viper.GetBool("oos.minio.useSSL") {
		link = fmt.Sprintf("https://%s/%s/%s", viper.GetString("oos.minio.addr"), bucket, url)
	} else {
		link = fmt.Sprintf("http://%s/%s/%s", viper.GetString("oos.minio.addr"), bucket, url)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"message":"Avatar uploaded successfully.",
		"url": link,
	})
}
