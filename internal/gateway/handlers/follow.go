package handlers

import (
	"github.com/gin-gonic/gin"
)

func FollowReqHandler(c *gin.Context) {
	//name := middleware.GetUserName(c)
	//target := c.PostForm("following")
	//
	//cli, conn, err := client.Accounts()
	//if err != nil {
	//	log.Println(err)
	//}
	//defer conn.Close()
	//
	//r, err := cli.NewFollows(context.Background(), &pb.FollowersData{
	//	Follower:  name,
	//	Following: target,
	//})
	//if err != nil {
	//	log.Printf("Failed to send message to Accounts server: %v", err)
	//}
	//
	//
	//c.JSON(200, r)

}

//func FollowAcceptHandler(c *gin.Context) {
//	inboxID := c.Param("id")
//
//
//}