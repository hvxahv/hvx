package cache

//// DELKey Delete the redis key by passing in the key.
//func DELKey(k string) error {
//	rd := GetRDB()
//	if err := rd.Del(ctx, k).Err(); err != nil {
//		return errors.New("failed to delete redis key.")
//	}
//
//	return nil
//}
//
//// DELAcctMail Delete mail in the collection.
//func DELAcctMail(mail string) error {
//	ok := SISAcctMail(mail)
//	if !ok {
//		return errors.New("mail does't exist!")
//	}
//
//	err := DELSet("ACCT_MAIL", mail)
//	if err != nil {
//		return err
//	}
//
//	return nil
//
//}
//
//// DELSet Delete the value in the redis set and receive two parameters, key and member.
//func DELSet(key, member string) error {
//	rd := GetRDB()
//	err := rd.SRem(ctx, key, member).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
