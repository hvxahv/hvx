package accounts

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Aliases struct {
	gorm.Model
	Username string `gorm:"primaryKey;type:varchar(100);username;" validate:"required,min=10,max=16"`
}


func NewAliases(username string) (*Aliases, error) {
	v := validator.New()

	aa := &Aliases{Username: username}
	if err := v.Struct(*aa); err != nil {
		return nil, err
	}
	return aa, nil
}

type Index interface {
	// GetAcctAliases Get user ID by username.
	GetAcctAliases() (uint, error)
}

func (a *Aliases) GetAcctAliases() (uint, error) {
	id, err := FindAliases(a.Username)
	if err != nil {
		return 0, err
	}
	return id, nil
}


