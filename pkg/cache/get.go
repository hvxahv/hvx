package cache

func GetAccount(username string) (string, error) {
	r, err := GetRDB().Get(ctx, username).Result()
	if err != nil {
		return "", err
	}
	return r, nil
}
