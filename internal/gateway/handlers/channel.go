package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/channels"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/pkg/microservices/client"
	"golang.org/x/net/context"
)

func GetManagedChannelsHandler(c *gin.Context) {
	name := middleware.GetUsername(c)

	acct, err := accounts.NewAccountsUsername(name).GetAccountByUsername()
	if err != nil {
		return
	}
	ch, err := channels.NewChannelsOwnerID(acct.ID).GetByOwnerID()
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": ch,
	})
}

func CreateChannelHandler(c *gin.Context) {
	name := middleware.GetUsername(c)

	id := c.PostForm("id")
	cn := c.PostForm("name")
	bio := c.PostForm("bio")
	avatar := c.PostForm("avatar")

	isPrivate, err := strconv.ParseBool(c.PostForm("is_private"))
	if err != nil {
		return
	}
	if err := channels.NewChannels(cn, id, avatar, bio, name, isPrivate).Create(); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
	})
}

func UpdateChannelHandler(c *gin.Context) {

}

func DeleteChannelHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	id, err := strconv.Atoi(c.PostForm("id"))

	if err != nil {
		return
	}
	if err := channels.NewDeleteChannelByID(name, uint(id)).Delete(); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
	})
}

func NewChannelAdminHandler(c *gin.Context) {
	name := middleware.GetUsername(c)

	id := c.PostForm("id")
	admin := c.PostForm("admin")

	cli, conn, err := client.Channel()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewAdmin(context.Background(), &pb.NewAdminData{
		Owner: name,
		Id:    id,
		Admin: admin,
	})

	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})
}

func RemoveChannelAdminHandler(c *gin.Context) {
	name := middleware.GetUsername(c)

	fmt.Println(name)
}

func NewSubscriberHandler(c *gin.Context) {
	name := middleware.GetUsername(c)
	id := c.PostForm("id")

	cli, conn, err := client.Channel()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewSubscriber(context.Background(), &pb.NewSubscriberData{
		Id:   id,
		Name: name,
	})

	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})
}

func GetSubscriberListHandler(c *gin.Context) {

}
