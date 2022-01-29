package cache

func GetAccount(username string) (string, error) {
	r, err := GetRDB().Get(ctx, username).Result()
	if err != nil {
		return "", err
	}
	return r, nil
}

func GETDHData(deviceID string) ([]byte, error) {
	rd := GetRDB()
	d, err := rd.Get(ctx, deviceID).Bytes()
	if err != nil {
		return d, err
	}
	return d, nil
}
