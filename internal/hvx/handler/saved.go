package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"github.com/hvxahv/hvxahv/internal/saved"
)

func SavedHandler(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	hash := c.PostForm("hash")
	types := c.PostForm("types")
	client, err := saved.NewSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.NewSavedCreate{
		AccountId:   middleware.GetAccountID(c),
		Name:        name,
		Description: description,
		Hash:        hash,
		Types:       types,
	}

	create, err := client.Create(c, d)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"code":    create.Code,
		"message": create.Reply,
	})
}

func GetSavedHandler(c *gin.Context) {
	client, err := saved.NewSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.NewSavedID{Id: c.Param("id")}
	s, err := client.GetSaved(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"id":          s.Id,
		"name":        s.Name,
		"description": s.Description,
		"hash":        s.Hash,
		"types":       s.Types,
		"created_at":  s.CreatedAt,
	})
}

func GetSavesHandler(c *gin.Context) {
	client, err := saved.NewSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.NewSavedAccountID{
		AccountId: middleware.GetAccountID(c),
	}
	s, err := client.GetSaves(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  s.Code,
		"saves": s.Saves,
	})
}
