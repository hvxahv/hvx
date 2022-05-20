package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PasswordVerificationFailed = "PASSWORD_VERIFICATION_FAILED"
)

type auth interface {
	Verify(password string) (*Accounts, error)
}

func NewVerify(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

// Verify ...
func (a *Accounts) Verify(password string) (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).Where("username = ?", a.Username).First(&a).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return nil, errors.Errorf(PasswordVerificationFailed)
	}

	return &Accounts{
		Model: gorm.Model{
			ID: a.ID,
		},
		Mail: a.Mail,
	}, nil
}
