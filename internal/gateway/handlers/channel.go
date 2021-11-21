package handlers

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

func NewChannelHandler(c *gin.Context) {
	name := middleware.GetUserName(c)

	id := c.PostForm("id")
	cn := c.PostForm("name")
	bio := c.PostForm("bio")
	avatar := c.PostForm("avatar")
	//is_private := c.PostForm("is_private")

	cli, conn, err := client.Channel()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.New(context.Background(), &pb.NewChannelData{
		Id:        id,
		Name:      cn,
		Bio:       bio,
		Avatar:    avatar,
		Owner:     name,
		IsPrivate: false,
	})

	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})
}

func UpdateChannelHandler(c *gin.Context) {

}

func DeleteChannelHandler(c *gin.Context) {

}

func NewChannelAdminHandler(c *gin.Context) {
	name := middleware.GetUserName(c)

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
	name := middleware.GetUserName(c)

	fmt.Println(name)
}

func NewSubscriberHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
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
