package handlers

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

func NewChannelHandler(c *gin.Context) {
	name, err := middleware.GetUserName(c)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	}

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

	r, err := cli.NewChannel(context.Background(), &pb.NewChannelData{
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

func NewChannelAdminHandler(c *gin.Context) {
	name, err := middleware.GetUserName(c)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err,
		})
	}

	id := c.PostForm("id")
	admin := c.PostForm("admin")

	cli, conn, err := client.Channel()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewChannelAdmin(context.Background(), &pb.NewChanAdmData{
		Owner: name,
		Id:    id,
		Admin: admin,
	})

	fmt.Println(r)
	c.JSON(int(r.Code), gin.H{
		"code":    r.Code,
		"message": r.Message,
	})
}
