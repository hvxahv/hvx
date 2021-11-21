package accounts

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/security"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Accounts struct {
	gorm.Model

	Username string `gorm:"primaryKey;type:text;preferredUsername;" validate:"required,min=4,max=16"`
	Mail     string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// When creating an account, first verify the username, email address, and password.
	// After the verification is successful, store the username and key in the actors table,
	// then use the returned ActorID in this field, and then store the data in the accounts table .
	// At this time, the context of creating the user is complete.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// Whether to set as a private account
	IsPrivate  bool   `gorm:"type:boolean;is_private"`
	PrivateKey string `gorm:"type:text;private_key"`
}

func (a *Accounts) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return err
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.ActorID).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return err
	}

	return nil
}

func NewAcctNameANDActorID(username string, id uint) *Accounts {
	return &Accounts{
		Username: username,
		ActorID:  id,
	}
}

func (a *Accounts) UpdateUsername(target string) error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).Update("username", target).Error; err != nil {
		return err
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.ActorID).Update("preferred_username", target).Error; err != nil {
		return err
	}

	return nil
}

func (a *Accounts) Update() error {
	if a.Password != "" {
		a.Password = security.GenPassword(a.Password)
	}

	db := cockroach.GetDB()

	err := db.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func NewAccountsName(username string) *Accounts {
	return &Accounts{Username: username}
}

func (a *Accounts) FindAccountByUsername() (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ? ", a.Username).First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func NewAccounts(username string, mail string, password string) *Accounts {
	hash := security.GenPassword(password)
	return &Accounts{Username: username, Mail: mail, Password: hash}
}

func (a *Accounts) Create() error {
	if err := validator.New().Struct(*a); err != nil {
		return err
	}

	// Before creating, first, check whether the user exists. If it does not exist, create the user.
	// If the account is found, it returns the user has created.
	// It will not be judged so detailed in the database. It only returns the information created by the user.
	// If more processing is required, detailed operations need to be done in the cache.

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return errors.New("FAILED_TO_AUTOMATICALLY_CREATE_DATABASE")
	}

	if err := db.Debug().Table("accounts").Where("username = ? ", a.Username).Or("mail = ?", a.Mail).First(&Accounts{}).Error; err != nil {
		ok := cockroach.IsNotFound(err)
		if !ok {
			return errors.New("THE_ACCOUNT_ALREADY_EXISTS")
		}
	}

	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("failed to generate public and private keys: %v", err)
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	acct, err := NewActors(a.Username, a.Password, publicKey, "Person").NewActor()
	if err != nil {
		return err
	}

	a.ActorID = acct.ID
	a.PrivateKey = privateKey

	if err := db.Debug().Table("accounts").Create(&a).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	return nil
}

type Account interface {

	Create() error

	FindAccountByUsername() (*Accounts, error)

	Update() error

	// UpdateUsername Change the username and pass the target username as a parameter.
	UpdateUsername(target string) error

	Delete() error
}
