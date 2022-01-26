package account

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Authorization The interface for account authorization,
// including method interfaces for login or developer API permissions.
type Authorization interface {
	// VerifyAccount to the account.
	VerifyAccount() (*VerifyAccountReply, error)
}

func NewAuth(username string, password string) Authorization {
	return &Accounts{Username: username, Password: password}
}

type VerifyAccountReply struct {
	AccountID uint
	Mail      string
	Username  string
	Password  string
}

func (a *Accounts) VerifyAccount() (*VerifyAccountReply, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).First(&acct).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acct.Password), []byte(a.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	return &VerifyAccountReply{
		AccountID: acct.ID,
		Mail:      acct.Mail,
		Username:  acct.Username,
		Password:  acct.Password,
	}, nil
}
