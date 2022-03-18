/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/saved"
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
)

func GetSaves(c *gin.Context) {
	id := middleware.GetAccountID(c)
	client, err := saved.NewSavedClient()
	if err != nil {
		return
	}
	d := &v1alpha1.GetSavesRequest{AccountId: id}
	saves, err := client.GetSaves(c, d)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"saves": saves,
	})
}
