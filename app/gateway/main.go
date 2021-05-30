/*
整个程序的入口，提供外部访问程序的唯一入口
内部服务只能通过此服务才能进行调用
提供：
1. HTTP REST API 接口服务
2. 鉴权
*/
package main

import (
	"hvxahv/pkg/config"
	"hvxahv/pkg/maria"
	"hvxahv/pkg/mongo"
	"hvxahv/pkg/redis"
	"log"
)

func main() {
	if err := config.InitConfig("local"); err != nil {
		log.Println(err)
	}
	if err := maria.InitMariaDB(); err != nil {
		log.Println(err)
	}
	if err := mongo.InitMongoDB(); err != nil {
		log.Println(err)
	}

	if err := redis.InitRedis(); err != nil {
		log.Println(err)
	}

	//r := server.IngressRouter()
	//go bot.ServicesRunningNotice("gateway gateway", "7000")
	//_ = r.Run(fmt.Sprintf(":%s", viper.GetString("port.gateway")))
}
