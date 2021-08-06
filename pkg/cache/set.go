package cache

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

// SETAcct Set up the account cache function,
//If the key of the account name exists in the cache, delete it,
// and then add the data to the cache. If it does not exist,
// add the data to the cache directly.
func SETAcct(k string, v []byte, exp time.Duration) error {
	rd := GetRDB()
	if ok := ExistAcct(k); ok {
		if err := DelKey(k); err != nil {
			return err
		}
	}
	if err := rd.Set(ctx, k, v, exp).Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


// SETAcctMailORUN After successful registration,
// put the user's email and username in the cache collection.
func SETAcctMailORUN(mail, user string) error {
	rd := GetRDB()
	if err := rd.SAdd(ctx, "ACCT_MAIL", mail).Err(); err != nil {
		return errors.Errorf("failed to store mail in cache collection.")
	}

	if err := rd.SAdd(ctx, "ACCT_USERNAME", user).Err(); err != nil {
		return errors.Errorf("failed to store username in cache collection.")
	}
	return nil
}
