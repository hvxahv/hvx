/*
整个程序的入口，提供外部访问程序的唯一入口
内部服务只能通过此服务才能进行调用
提供：
1. HTTP REST API 接口服务
2. 鉴权
3. 访问微服务的 GRPC 客户端
*/
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/database"
	"log"
)

func main()  {

	if err := database.InitMariaDB(); err != nil {
		log.Println(err)
	}
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r := IngressRouter()
	go bot.ServicesRunningNotice("gateway gateway", "7000")
	_ = r.Run(fmt.Sprintf(":%s", viper.GetString("port.gateway")))
}

