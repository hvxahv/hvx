package accounts

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AccountAliases struct {
	gorm.Model
	Username string `gorm:"primaryKey;type:varchar(100);username;" validate:"required,min=10,max=16"`
}


func NewAccountID(username string) (*AccountAliases, error) {
	v := validator.New()

	aa := &AccountAliases{Username: username}
	if err := v.Struct(*aa); err != nil {
		return nil, err
	}
	return aa, nil
}

type Index interface {
	// GetAcctID Get user ID by username.
	GetAcctID() (uint, error)
}

func (a *AccountAliases) GetAcctID() (uint, error) {
	id, err := FindAccountID(a.Username)
	if err != nil {
		return 0, err
	}
	return id, nil
}


