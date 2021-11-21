package handlers

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/articles"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/pkg/microservices/client"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

func NewArticleHandler(c *gin.Context) {
	// Use this client to remotely call the login method.
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	r, err := cli.FindAccountsByUsername(context.Background(), &pb.AccountUsername{
		Username: middleware.GetUserName(c),
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	title := c.PostForm("title")
	summary := c.PostForm("summary")
	article := c.PostForm("article")
	nsfw := c.PostForm("nsfw")
	isArticle := c.PostForm("isArticle")


	isA, err := strconv.ParseBool(isArticle)
	if err != nil {
		log.Println(err)
	}

	isNSFW, err := strconv.ParseBool(nsfw)
	if err != nil {
		log.Println(err)
	}

	if isA {
		if err := articles.NewArticles(uint(r.Id), r.Username, title, summary, article, isNSFW).Create(); err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println("创建状态")
		if err := articles.NewStatus(uint(r.Id), r.Username, article, isNSFW).Create(); err != nil {
			log.Println(err)
		}
	}
}

func GetArticleHandler(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	n := articles.NewArticleID(uint(intID))
	article, err := n.FindByID()
	if err != nil {
		log.Println(err)
		return 
	}

	c.JSON(200, gin.H{
		"id": article.ID,
		"name": article.AuthorName,
		"title": article.Title,
		"summary": article.Summary,
		"article": article.Article,
		"nsfw": article.NSFW,
		"statuses": article.Statuses,
	})
}
