/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateArticleHandler(c *gin.Context) {
	to := c.PostFormArray("to")
	cc := c.PostFormArray("cc")
	client, err := article.GetArticleClient()
	if err != nil {
		return
	}
	articles, err := client.CreateArticle(c, &pb.CreateArticleRequest{
		AccountId:      middleware.GetAccountID(c),
		Title:          c.PostForm("title"),
		Summary:        c.PostForm("summary"),
		Article:        c.PostForm("article"),
		Tags:           c.PostFormArray("tags"),
		AttachmentType: c.PostForm("attachmentType"),
		Attachments:    c.PostFormArray("attachments"),
		To:             to,
		Cc:             cc,
		State:          c.GetBool("state"),
		Nsfw:           c.GetBool("nsfw"),
		Visibility:     c.PostForm("visibility"),
	})

	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": articles.Reply,
	})
}

func UpdateArticleHandler(c *gin.Context) {
	client, err := article.GetArticleClient()
	if err != nil {
		return
	}
	articles, err := client.UpdateArticle(c, &pb.UpdateArticleRequest{
		Id:             c.PostForm("id"),
		AccountId:      middleware.GetAccountID(c),
		Title:          c.PostForm("title"),
		Summary:        c.PostForm("summary"),
		Article:        c.PostForm("article"),
		Tags:           c.PostFormArray("tags"),
		AttachmentType: c.PostForm("attachmentType"),
		Attachments:    c.PostFormArray("attachments"),
		Nsfw:           c.GetBool("nsfw"),
		Visibility:     c.PostForm("visibility"),
	})

	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": articles.Reply,
	})
}

func GetArticleHandler(c *gin.Context) {
	client, err := article.GetArticleClient()
	if err != nil {
		return
	}
	articles, err := client.GetArticle(c, &pb.GetArticleRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		return
	}
	c.JSON(200, articles)
}

func GetArticlesHandler(c *gin.Context) {
	client, err := article.GetArticleClient()
	if err != nil {
		return
	}
	articles, err := client.GetArticlesByAccountID(c, &pb.GetArticlesByAccountIDRequest{
		AccountId: middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, articles)
}

func DeleteArticleHandler(c *gin.Context) {
	id := c.Param("id")
	client, err := article.GetArticleClient()
	if err != nil {
		return
	}
	articles, err := client.DeleteArticle(c, &pb.DeleteArticleRequest{
		Id:        id,
		AccountId: middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, articles)
}
