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

func DeleteAccount(c *gin.Context) {
	username := middleware.GetUsername(c)
	password := c.PostForm("password")
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewAccountDelete{Username: username, Password: password}
	r, err := client.Delete(c, d)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"reply": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": r.Reply,
	})
}

func EditAccountUsernameHandler(c *gin.Context) {
	username := c.PostForm("username")
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewEditAccountUsername{
		Id:       middleware.GetAccountID(c),
		Username: username,
	}
	r, err := client.EditUsername(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": r.Reply,
	})
}

func EditAccountPasswordHandler(c *gin.Context) {
	password := c.PostForm("password")
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewEditAccountPassword{
		Id:       middleware.GetAccountID(c),
		Password: password,
	}
	r, err := client.EditPassword(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": r.Reply,
	})
}

func EditAccountMailHandler(c *gin.Context) {
	mail := c.PostForm("mail")
	client, err := account.NewAccountClient()
	if err != nil {
		return
	}
	d := &pb.NewEditAccountMail{Id: middleware.GetAccountID(c), Mail: mail}
	edit, err := client.EditMail(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": edit.Reply,
	})
}
