package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/storage"
	"github.com/spf13/viper"
)

func SignUpHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	mail := c.PostForm("mail")
	publicKey := c.PostForm("publicKey")

	// Use the client to call the Accounts service to create users.
	//cli, conn, err := client.Accounts()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer conn.Close()
	//
	//r, err := cli.Create(context.Background(), &pb.CreateAccountData{
	//	Username: username,
	//	Password: password,
	//	Mail:     mail,
	//})
	//if err != nil {
	//	log.Printf("failed to send message to accounts server: %v", err)
	//}

	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	actor, err := accounts.NewActors(username, publicKey, "Person").Create()
	if err != nil {
		log.Println(err)
		return
	}

	if err := accounts.NewAccounts(username, mail, password, actor.ID).Create(); err != nil {
		log.Panicln(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok",
	})

}

// UploadAvatar Interface for users to upload avatars.
func UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("File read failed！" + err.Error())
		c.JSON(500, gin.H{
			"status":  500,
			"message": "file read failed！",
		})
		return
	}

	fileType := file.Header.Get("Content-Type")
	if fileType != "image/jpeg" && fileType != "image/png" {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "unsupported image format!",
		})
		return
	}

	bucket := "avatar"
	m := storage.NewMinio(file, bucket, viper.GetString("minio.location"), fileType)
	url, err := m.Uploader()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "server receiving file error!",
		})
		return
	}

	var link string
	if viper.GetBool("storage.minio.useSSL") {
		link = fmt.Sprintf("https://%s/%s/%s", viper.GetString("minio.addr"), bucket, url)
	} else {
		link = fmt.Sprintf("http://%s/%s/%s", viper.GetString("minio.addr"), bucket, url)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "avatar uploaded successfully.",
		"url":     link,
	})
}

// FetchAccountsHandler Obtain personal account information,
// analyze the user through TOKEN and return user data.
//func GetAccountsHandler(c *gin.Context) {
//	name := middleware.GetUserName(c)
//
//	cli, conn, err := client.Accounts()
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//
//	r, err := cli.QueryByName(context.Background(), &pb.NewAccountByName{
//		Username: name,
//	})
//	if err != nil {
//		log.Printf("failed to send message to accounts server: %v", err)
//	}
//
//	c.JSON(200, gin.H{
//		"code":     200,
//		"activity": r,
//	})
//}

//func DeleteAccount(c *gin.Context) {
//	password := c.PostForm("password")
//	mail := c.PostForm("mail")
//
//	cli, conn, err := client.Accounts()
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//
//	r, err := cli.Delete(context.Background(), &pb.AuthData{
//		Mail:     mail,
//		Password: password,
//	})
//	if err != nil {
//		log.Printf("failed to send message to accounts server: %v", err)
//	}
//
//	c.JSON(int(r.Code), gin.H{
//		"code":    r.Code,
//		"message": r.Message,
//	})
//
//}

//func UpdateAccount(c *gin.Context) {
//	username := middleware.GetUserName(c)
//	password := c.PostForm("password")
//	bio := c.PostForm("bio")
//	name := c.PostForm("name")
//	mail := c.PostForm("mail")
//	phone := c.PostForm("phone")
//	is_private := c.PostForm("is_private")
//
//	cli, conn, err := client.Accounts()
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//
//	parseBool, err := strconv.ParseBool(is_private)
//	if err != nil {
//		log.Println(err)
//	}
//
//	r, err := cli.Update(context.Background(), &pb.AccountData{
//		Username:  username,
//		Password:  password,
//		Mail:      mail,
//		Bio:       bio,
//		Name:      name,
//		Phone:     phone,
//		IsPrivate: parseBool,
//	})
//
//	if err != nil {
//		log.Printf("failed to send message to accounts server: %v", err)
//	}
//
//	c.JSON(int(r.Code), gin.H{
//		"code":    r.Code,
//		"message": r.Message,
//	})
//
//}
