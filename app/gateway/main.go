/*
整个程序的入口网关，提供外部访问程序的唯一入口
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
	"hvxahv/app/gateway/pkg/social"
	"hvxahv/app/gateway/pkg/tools"
	"hvxahv/pkg/auth"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/http"
)

func main()  {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r := gin.Default()
	r.Use(http.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* 账号登录和注册 */
	r.POST("/account/new", NewAccountsHandler)
	r.POST("/account/login", tools.AuthHandler)

	// 通过 Token 才能访问的功能
	v1 := r.Group("/api/v1")
	v1.Use(auth.JWTAuth)
	{
		/* Accounts Services */
		v1.GET("/account/i", GetAccountsHandler)
		//v1.POST("/account.bac/delete", account.bac.DeleteAccountHandler)
		//v1.POST("/account.bac/settings", account.bac.AccountSettingHandler)

		/*  Article Services */
		v1.POST("/article/new", social.CreateArticleHandler)
		v1.POST("/article/update", social.UpdateArticleHandler)
		v1.POST("/article/delete", social.DeleteArticleHandler)

		/* Status Services */
		v1.POST("/status/new", social.CreateStatusHandler)
		v1.POST("/status/update", social.UpdateStatusListHandler)
		v1.POST("/status/delete", social.DeleteStatusHandler)
	}

	go bot.ServicesRunningNotice("ingress gateway", "7000")
	r.Run(":7000")
}

