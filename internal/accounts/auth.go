package accounts

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// AccountAuth The interface for account authorization,
// including method interfaces for login or developer API permissions.
type AccountAuth interface {
	// Login to the account and generate token, Return token and custom error message.
	Login() (string, error)

}

func NewAccountAuth(mail string, password string) AccountAuth {
	return &Accounts{Mail: mail, Password: password}
}

func (a *Accounts) Login() (string, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	if err := db.Debug().Table("accounts").Where("mail = ?", a.Mail).First(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(a.Password)); err != nil {
		return "", errors.Errorf("Password verification failed.")
	}

	return acct.Name, nil
}
