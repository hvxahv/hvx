package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

// ExistAcct Check if redis exists, if KEY is available, return true, otherwise return false.
// https://redis.uptrace.dev/#redisnil
func ExistAcct(key string) bool {
	_, err := GetRDB().Get(ctx, key).Result()
	if err != redis.Nil {
		return true
	}
	return false
}

func SETAcctHash(key, field string, value interface{}) error {
	ad, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
		return err
	}

	data := make(map[string]interface{})
	data[field] = ad

	if err := GetRDB().HMSet(ctx, key, data).Err(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

