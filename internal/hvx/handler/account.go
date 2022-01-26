package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
)

// UploadAvatar Interface for users to upload avatars.
//func UploadAvatar(c *gin.Context) {
//	avatar, err := c.FormFile("avatar")
//	if err != nil {
//		fmt.Println("File read failed！" + err.Error())
//		c.JSON(500, gin.H{
//			"status":  500,
//			"message": "file read failed！",
//		})
//		return
//	}
//
//	fileType := avatar.Header.Get("Content-Type")
//	if fileType != "image/jpeg" && fileType != "image/png" {
//		c.JSON(500, gin.H{
//			"status":  500,
//			"message": "unsupported image format!",
//		})
//		return
//	}
//
//	file, err := avatar.Open()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	info, err := minio.NewFilesUploader("avatar", avatar.Filename, fileType, file).Uploader()
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	log.Println(info.Location)
//
//	actor, err := account.NewActorByAccountUsername(middleware.GetUsername(c)).GetActorByAccountUsername()
//	if err != nil {
//		return
//	}
//	a := account.NewActorID(actor.ID)
//	a.Avatar = info.Location
//	if err := a.Update(); err != nil {
//		log.Println(err)
//		return
//	}
//	c.JSON(200, gin.H{
//		"status":  200,
//		"message": "avatar uploaded successfully.",
//		"url":     info.Location,
//	})
//}

//GetAccountHandler Obtain personal account information,
//analyze the user through TOKEN and return user data.
func GetAccountHandler(c *gin.Context) {
	client, err := account.NewClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountUsername{Username: middleware.GetUsername(c)}
	actor, err := client.GetActorByAccountUsername(c, d)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":  200,
		"actor": actor,
	})
}

func DeleteAccount(c *gin.Context) {
	username := middleware.GetUsername(c)
	password := c.PostForm("password")
	client, err := account.NewClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountDelete{Username: username, Password: password}
	da, err := client.DeleteAccount(c, d)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "500",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": da.Reply,
	})
}

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
//		log.Printf("failed to send message to account server: %v", err)
//	}
//
//	c.JSON(int(r.Code), gin.H{
//		"code":    r.Code,
//		"message": r.Message,
//	})
//
//}
