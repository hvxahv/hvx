package redis

import (
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// ExistKey Check if redis exists, if KEY is available, return true, otherwise return false.
// https://redis.uptrace.dev/#redisnil
func ExistKey(key string) bool {
	_, err := GetRDB().Get(ctx, key).Result()
	if err != redis.Nil {
		return true
	}
	return false
}

func SetJsonData(key string, value []byte, expiration time.Duration) error {
	_, err := GetRDB().Set(ctx, key, value, expiration).Result()
	if err != nil {
		log.Println("failed to store to cache:", err)
		return err
	}
	return nil
}