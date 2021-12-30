package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/internal/channels"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"log"
	"strconv"
)

func CreateBroadcastHandler(c *gin.Context) {
	channelID := c.PostForm("channel_id")
	a, err := accounts.NewAccountsUsername(middleware.GetUsername(c)).GetAccountByUsername()
	if err != nil {
		log.Println(err)
		return
	}
	title := c.PostForm("title")
	summary := c.PostForm("summary")
	article := c.PostForm("article")
	nsfw := c.PostForm("nsfw")

	ci, err := strconv.Atoi(channelID)
	if err != nil {
		log.Println(err)
		return
	}
	n, err := strconv.ParseBool(nsfw)
	if err != nil {
		log.Println(err)
		return
	}

	cid, err := channels.NewBroadcastsIPFSCID(uint(ci), a.ActorID, title, summary, article, n)
	if err != nil {
		log.Println(err)
		return
	}
	if err := channels.NewBroadcasts(uint(ci), a.ActorID, cid).Create(); err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}

func GetBroadcastsHandler(c *gin.Context) {
	ci, err := strconv.Atoi(c.Query("channel_id"))
	if err != nil {
		log.Println(err)
		return
	}
	b, err := channels.NewBroadcastsChannelID(uint(ci)).GetBroadcastsByChannelID()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":       "200",
		"broadcasts": b,
	})
}
