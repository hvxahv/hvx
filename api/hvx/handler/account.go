package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
)

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
