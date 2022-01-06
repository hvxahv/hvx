package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// GenPassword Use the bcrypt package to crypto the password and return the encrypted hash,
// which needs to be converted into a string.
func GenPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Encryption password failed: ", err)
	}
	return string(hash)
}
