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

func GetSaves(c *gin.Context) {
	id := middleware.GetAccountID(c)
	client, err := saved.GetSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.GetSavesRequest{AccountId: id}
	saves, err := client.GetSaves(c, d)
	if err != nil {
		return
	}
	c.JSON(200, saves)
}

func GetSaved(c *gin.Context) {
	aid := middleware.GetAccountID(c)
	id := c.Param("id")

	client, err := saved.GetSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.GetSavedRequest{
		AccountId: aid,
		Id:        id,
	}
	save, err := client.GetSaved(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"saved": save,
	})
}

func CreateSaved(c *gin.Context) {
	aid := middleware.GetAccountID(c)
	name := c.PostForm("name")
	description := c.PostForm("description")
	cid := c.PostForm("cid")
	types := c.PostForm("types")

	client, err := saved.GetSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.CreateSavedRequest{
		AccountId:   aid,
		Name:        name,
		Description: description,
		Cid:         cid,
		Types:       types,
		IsPrivate:   false,
	}
	save, err := client.CreateSaved(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": save.Reply,
	})
}

func EditSaved(c *gin.Context) {
	aid := middleware.GetAccountID(c)
	id := c.Param("id")
	name := c.PostForm("name")
	description := c.PostForm("description")

	client, err := saved.GetSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.EditSavedRequest{
		Id:          id,
		AccountId:   aid,
		Name:        name,
		Description: description,
	}
	save, err := client.EditSaved(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": save.Reply,
	})
}

func DeleteSaved(c *gin.Context) {
	aid := middleware.GetAccountID(c)
	id := c.Param("id")

	client, err := saved.GetSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.DeleteSavedRequest{
		Id:        id,
		AccountId: aid,
	}
	save, err := client.DeleteSaved(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": save.Reply,
	})
}
