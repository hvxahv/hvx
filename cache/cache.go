package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/hvxahv/hvx/errors"
	"golang.org/x/net/context"
)

type Cache struct {
	Ctx    context.Context
	Client *redis.Client
	Err    error
}

func NewCache(db int) *Cache {
	if err := NewRdb().Dial(db); err != nil {
		errors.Throw("CONNECT REDIS ERRORS", err)
		return &Cache{Err: err}
	}
	return &Cache{Ctx: context.Background(), Client: GetRDB(), Err: nil}
}

type Operator interface {
	// String operations.
	//EXISTS(key string) (bool, error)
	//GET(key string) (string, error)

	SETDH(deviceId string, data []byte) error
	GETDH(deviceId string) ([]byte, error)
	//DEL(key string) error

	// Hash operations.
	//HEXISTS(key string, field string) (bool, error)
	//HSET(key string, field string, value string) error
	//HGET(key string, field string) (string, error)
	//HGETALL(key string) (map[string]string, error)
	//HMSET(key string, fields map[string]string) error
	//HMGET(key string, fields []string) ([]string, error)
}
