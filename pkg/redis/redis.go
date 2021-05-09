package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var (
	pool *redis.Pool
)
func InitRedis() error {
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
				panic(fmt.Sprintf("Can't connect to redis %s address: %s", err, addr))
			}
			return c, err
		},
	}
	return err
}

func GetRDB() redis.Conn {
	return pool.Get()
}
