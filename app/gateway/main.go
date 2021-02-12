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
	"github.com/spf13/viper"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/db"
	"log"
)

func main()  {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := db.InitMariaDB(); err != nil {
		log.Println(err)
	}
	if err := db.InitMongoDB(); err != nil {
		log.Println(err)
	}

	db.InitRedis()

	r := IngressRouter()
	go bot.ServicesRunningNotice("gateway gateway", "7000")
	_ = r.Run(fmt.Sprintf(":%s", viper.GetString("port.gateway")))
}

