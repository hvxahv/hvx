/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/spf13/viper"
	"log"
)

func GetPublicAccountCountHandler(c *gin.Context) {
	db := cockroach.GetDB()
	var count int64
	if err := db.Debug().Table("accounts").Count(&count).Error; err != nil {
		log.Println(err)
		return
	}
	fmt.Println(count)
	c.JSON(200, gin.H{
		"code":          "200",
		"account_count": count,
	})
}

func GetPublicInstanceDetailsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "200",
		"details": gin.H{
			"version":    viper.GetString("version"),
			"build":      "2022-01-01",
			"maintainer": viper.GetString("author"),
			"repo":       viper.GetString("name"),
			"host":       viper.GetString("localhost"),
		},
	})
}

func GetNodeInfoHandler(c *gin.Context) {

}
