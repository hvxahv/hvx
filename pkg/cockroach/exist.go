package cockroach

import (
	"gorm.io/gorm"
)

// IsNotFound If the error is NOT FOUND, return TRUE otherwise return FALSE.
func IsNotFound(err error) (bool, error) {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	return false, nil
}