package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"log"
)

var (
	pool *redis.Pool
)
func InitRedis() {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	addr := viper.GetString("redis.address")

	pwd := redis.DialPassword(viper.GetString("redis.password"))
	pool = &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, pwd)
			if err != nil {
				log.Printf("连接不上 redis : %s 地址: %s", err, addr)
			}
			return c, err
		},
	}
}

func GetRDB() redis.Conn {
	return pool.Get()
}
