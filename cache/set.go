package cache

import (
	"time"
)

//
//// SETAcct Set up the account cache function,
//// If the key of the account name exists in the cache, delete it,
//// and then add the data to the cache. If it does not exist,
//// add the data to the cache directly.
//func SETAcct(k string, v []byte, exp time.Duration) error {
//	rd := GetRDB()
//	if ok := SISAcct(k); ok {
//		if err := DELKey(k); err != nil {
//			return err
//		}
//	}
//	if err := rd.Set(ctx, k, v, exp).Err(); err != nil {
//		fmt.Println(err)
//		return err
//	}
//	return nil
//}
//
//// SETAcctMail After successful registration,
//// put the email the cache collection.
//func SETAcctMail(mail string) error {
//	rd := GetRDB()
//	if err := rd.SAdd(ctx, "ACCT_MAIL", mail).Err(); err != nil {
//		return errors.New("failed to store mail in cache collection.")
//	}
//
//	return nil
//}

func (r *Cache) SETDH(deviceId string, data []byte) error {
	if err := r.Client.Set(r.Ctx, deviceId, data, 1200*time.Second).Err(); err != nil {
		return err
	}
	return nil
}
