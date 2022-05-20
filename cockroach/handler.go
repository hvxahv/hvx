package cockroach

import "gorm.io/gorm"

// IsNotFound If the error is NOT FOUND, return TRUE otherwise return FALSE.
func IsNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}
