package account

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// Authenticate The interface for account authorization,
// including method interfaces for login or developer API permissions.
type Authenticate interface {
	// SignIn to the account and generate token, Return token and custom error message.
	SignIn() (uint, string, error)
}

type Authenticates struct {
	Username string
	Password string
}

func NewAuthenticate(username string, password string) *Authenticates {
	return &Authenticates{Username: username, Password: password}
}

func (a *Authenticates) SignIn() (uint, string, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	if err := db.Debug().Table("account").Where("username = ?", a.Username).First(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return 0, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acct.Password), []byte(a.Password)); err != nil {
		return 0, "", errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	return acct.ID, acct.Mail, nil
}
