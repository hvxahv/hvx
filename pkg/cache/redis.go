package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var ctx = context.Background()

// rdb_Once is executed when the called, and it is executed only once.
var once sync.Once
var rdb *redis.Client

// InitRedis Initialize cache, set parameters, and return to cache client.
func InitRedis(db int) {
	h := viper.GetString("redis.host")
	p := viper.GetString("redis.port")
	pwd := viper.GetString("redis.password")
	addr := fmt.Sprintf("%s:%s", h, p)

	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Network:            "tcp",
			Addr:               addr,
			Dialer:             nil,
			OnConnect:          nil,
			Username:           "",
			Password:           pwd,
			DB:                 db,
			MaxRetries:         0,
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        5 * time.Second,
			ReadTimeout:        3 * time.Second,
			WriteTimeout:       3 * time.Second,
			PoolSize:           15,
			MinIdleConns:       10,
			MaxConnAge:         0,
			PoolTimeout:        4 * time.Second,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
			Limiter:            nil,
		})

		pong, err := rdb.Ping(ctx).Result()
		if err != nil {
			log.Fatal(fmt.Sprintf("Connection cache error: %s", err))
		}
		log.Println(pong)
	})
}

// GetRDB Get the initialized cache client.
func GetRDB() *redis.Client {
	return rdb
}
