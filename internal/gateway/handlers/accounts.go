package handlers

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"strconv"
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

	r, err := cli.New(context.Background(), &pb.NewAccountData{
		Username: username,
		Password: password,
		Mail:     mail,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

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
	r, err := cli.Login(context.Background(), &pb.AuthData{
		Mail:     mail,
		Password: password,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	t, err := security.GenToken(r.Uuid, r.Username, password)

	c.JSON(200, gin.H{
		"code":     "200",
		"username": r.Username,
		"token":    t,
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
//	m := storage.NewMin(file, "avatar", bucket, "")
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
//	if viper.GetBool("storage.minio.useSSL") {
//		link = fmt.Sprintf("https://%s/%s/%s", viper.GetString("storage.minio.addr"), bucket, url)
//	} else {
//		link = fmt.Sprintf("http://%s/%s/%s", viper.GetString("storage.minio.addr"), bucket, url)
//	}
//	c.JSON(200, gin.H{
//		"status":  200,
//		"message": "Avatar uploaded successfully.",
//		"url":     link,
//	})
//}
//

// GetAccountsHandler Obtain personal account information,
// analyze the user through TOKEN and return user data.
func GetAccountsHandler(c *gin.Context) {
	name := middleware.GetUserName(c)

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.Find(context.Background(), &pb.NewAccountByName{
		Username: name,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	c.JSON(200, gin.H{
		"code":     200,
		"messages": r,
	})
}

func DeleteAccount(c *gin.Context) {
	password := c.PostForm("password")
	mail := c.PostForm("mail")

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.Delete(context.Background(), &pb.AuthData{
		Mail:     mail,
		Password: password,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})

}

func UpdateAccount(c *gin.Context) {
	username := middleware.GetUserName(c)
	password := c.PostForm("password")
	bio := c.PostForm("bio")
	name := c.PostForm("name")
	mail := c.PostForm("mail")
	phone := c.PostForm("phone")
	is_private := c.PostForm("is_private")

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	parseBool, err := strconv.ParseBool(is_private)
	if err != nil {
		fmt.Println(err)
	}

	r, err := cli.Update(context.Background(), &pb.AccountData{
		Username:  username,
		Password:  password,
		Mail:      mail,
		Bio:       bio,
		Name:      name,
		Phone:     phone,
		IsPrivate: parseBool,
	})

	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})

}
