package handlers

import (
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

func FollowHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	target := c.PostForm("following")

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	r, err := cli.NewFollow(context.Background(), &pb.FollowersData{
		Follower:  name,
		Following: target,
	})
	if err != nil {
		log.Printf("Failed to send message to Accounts server: %v", err)
	}


	c.JSON(200, r)

}
