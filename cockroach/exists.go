package cockroach

import (
	"gorm.io/gorm"
)

// IsNotFound Returns true if not found.
func IsNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}
