package cache

import "github.com/pkg/errors"

// DelKey Delete the redis key by passing in the key.
func DelKey(k string) error {
	rd := GetRDB()
	if err := rd.Del(ctx, k).Err(); err != nil {
		return errors.Errorf("failed to delete redis key.")
	}
	return nil
}
