package database

import (
	"github.com/gomodule/redigo/redis"
)

func Rdb() *redis.Pool{
	addr := "49.232.89.84:6379"
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
