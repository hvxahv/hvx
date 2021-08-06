package cache

import (
	"github.com/go-redis/redis/v8"
	"log"
)

// ExistAcct Check if cache exists, if KEY is available, return true, otherwise return false.
// https://redis.uptrace.dev/#redisnil
func ExistAcct(key string) bool {
	_, err := GetRDB().Get(ctx, key).Result()
	if err != redis.Nil {
		return true
	}
	return false
}

// FINDAcctMail When registering, check whether the user’s email and user name exist.
func FINDAcctMail(mail string) bool {
	rd := GetRDB()
	m, err := rd.SIsMember(ctx, "ACCT_MAIL", mail).Result()
	if err != nil {
		log.Println("failed to find members of mail in cache collection.")
	}
	return m

}

