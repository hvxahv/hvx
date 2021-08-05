package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
	"time"
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

func SETAcct(k string, v []byte, exp time.Duration) error {
	if err := GetRDB().Set(ctx, k, v, exp).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// SETAcctMailORUN After successful registration,
// put the user's email and username in the redis collection.
func SETAcctMailORUN(mail, user string) error {
	rd := GetRDB()
	if err := rd.SAdd(ctx, "ACCT_MAIL", mail).Err(); err != nil {
		return errors.Errorf("failed to store mail in redis collection.")
	}

	if err := rd.SAdd(ctx, "ACCT_USERNAME", user).Err(); err != nil {
		return errors.Errorf("failed to store username in redis collection.")
	}
	return nil
}

// FINDAcctMailAndUN When registering, check whether the user’s email and user name exist.
func FINDAcctMailAndUN(mail, user string) (bool, bool) {
	rd := GetRDB()
	m, err := rd.SIsMember(ctx, "ACCT_MAIL", mail).Result()
	if err != nil {
		log.Println("failed to find members of mail in redis collection.")
	}

	u, err := rd.SIsMember(ctx, "ACCT_USERNAME", user).Result()
	if err != nil {
		log.Println("failed to find members of username in redis collection.")
	}

	return m, u

}
