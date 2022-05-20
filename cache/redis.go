package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"log"
	"sync"
	"time"
)

var (
	ctx  = context.Background()
	once sync.Once
	rdb  *redis.Client
)

type option struct {
	addr     string
	password string
}

func NewRdb() *option {
	return &option{
		addr:     viper.GetString("redis.addr"),
		password: viper.GetString("redis.password"),
	}
}

func (r *option) Dial(db int) error {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Network:            "tcp",
			Addr:               r.addr,
			Dialer:             nil,
			OnConnect:          nil,
			Username:           "",
			Password:           r.password,
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
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Println(pong)
	return nil
}

func GetRDB() *redis.Client {
	return rdb
}
