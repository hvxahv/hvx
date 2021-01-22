/*
整个程序的入口，提供外部访问程序的唯一入口
内部服务只能通过此服务才能进行调用
提供：
1. HTTP REST API 接口服务
2. 鉴权
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hvxahv/app/ingress/handler"
	"hvxahv/app/ingress/pkg"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/middleware"
)

func main()  {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* 账号登录和注册 */
	r.POST("/account/new", handler.NewAccountsHandler)
	r.POST("/account/login", pkg.VerificationHandler)

	// 通过 Token 才能访问的功能
	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWTAuth)
	{
		/* Accounts Services */
		v1.GET("/account/i", handler.GetAccountsHandler)
		v1.POST("/account.bac/delete", handler.DeleteAccountHandler)
		v1.POST("/account.bac/settings", handler.AccountSettingHandler)

		/*  Article Services */
		v1.POST("/article/new", handler.CreateArticleHandler)
		v1.POST("/article/update", handler.UpdateArticleHandler)
		v1.POST("/article/delete", handler.DeleteArticleHandler)

		/* Status Services */
		v1.POST("/status/new", handler.CreateStatusHandler)
		v1.POST("/status/update", handler.UpdateStatusHandler)
		v1.POST("/status/delete", handler.DeleteStatusHandler)
	}

	go bot.ServicesRunningNotice("ingress ingress", "7000")
	_ = r.Run(":7000")
}

