package accounts

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// AccountAuth The interface for account authorization,
// including method interfaces for login or developer API permissions.
type AccountAuth interface {
	// SignIn to the account and generate token, Return token and custom error message.
	SignIn() (uint, string, error)
}

func NewAuth(username string, password string) AccountAuth {
	return &Accounts{Username: username, Password: password}
}

func (a *Accounts) SignIn() (uint, string, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).First(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return 0, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acct.Password), []byte(a.Password)); err != nil {
		return 0, "", errors.Errorf("Password verification failed.")
	}

	return acct.ID, acct.Mail, nil
}
