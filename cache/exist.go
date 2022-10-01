package cache

//// SISAcct Check if cache exists, if KEY is available, return true, otherwise return false.
//// If it returns true, it means that the username exists.
//// https://redis.uptrace.dev/#redisnil
//func SISAcct(key string) bool {
//	_, err := GetRDB().Get(ctx, key).Result()
//	if err != redis.Nil {
//		return true
//	}
//	return false
//}
//
//// SISAcctMail Check if the mail exist from the cache
//func SISAcctMail(mail string) bool {
//	rd := GetRDB()
//	m, err := rd.SIsMember(ctx, "ACCT_MAIL", mail).Result()
//	if err != nil {
//		log.Println("failed to find members of mail in cache collection.")
//	}
//	return m
//
//}
