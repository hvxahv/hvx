package account

import (
	"github.com/go-playground/validator/v10"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model

	Username  string `gorm:"primaryKey;type:text;preferredUsername;" validate:"required,min=4,max=16"`
	Mail      string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password  string `gorm:"type:text;password" validate:"required,min=8,max=100"`
	ActorID   uint   `gorm:"type:bigint;actor_id"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

func (a *Accounts) SetAccountPassword(password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	a.Password = string(hash)
	return a
}

func (a *Accounts) SetAccountUsername(username string) *Accounts {
	a.Username = username
	return a
}

func (a *Accounts) Create() error {
	if err := validator.New().Struct(*a); err != nil {
		return err
	}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACCOUNT_DATABASE")
	}

	if err := db.Debug().Table("accounts").Where("username = ? ", a.Username).Or("mail = ?", a.Mail).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.New("THE_USERNAME_OR_MAIL_ALREADY_EXISTS")
		}
	}

	if err := db.Debug().Table("accounts").Create(&a).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	return nil
}

func (a *Accounts) Update() error {
	if a.Username != "" {
		return errors.New("PLEASE_USE_THE_SET_ACCOUNT_PASSWORD_METHOD_TO_UPDATE_THE_USERNAME")
	}
	db := cockroach.GetDB()
	err := db.Debug().Table("accounts").Where("id = ?", a.ID).Updates(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Accounts) EditPassword() error {
	if err := NewAccountsID(a.ID).Update(); err != nil {
		return err
	}

	// Because the password is reset, all logged-in devices should be deleted
	if err := device.NewDevicesByAccountID(a.ID).DeleteAll(); err != nil {
		return err
	}
	return nil
}

func (a *Accounts) EditUsername() error {
	db := cockroach.GetDB()

	var acct Accounts
	if err := db.Debug().Table("accounts").Where("id = ?", a.ID).First(&acct).Update("username", a.Username).Error; err != nil {
		return err
	}
	if err := NewActorsID(acct.ActorID).SetActorPreferredUsername(a.Username).EditActorPreferredUsername(); err != nil {
		return err
	}

	return nil
}

func (a *Accounts) Delete() error {
	db := cockroach.GetDB()

	var acct Accounts
	if err := db.Debug().Table("accounts").Where("id = ?", a.ID).First(&acct).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return err
	}

	if err := NewActorsID(acct.ActorID).Delete(); err != nil {
		return err
	}

	return nil
}

func (a *Accounts) GetAccountByID() (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Accounts) GetAccountByUsername() (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func NewAccountsID(id uint) *Accounts {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewAccountsUsername(username string) *Accounts {
	return &Accounts{Username: username}
}

func NewAccounts(username, mail, password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: string(hash),
	}
}

type Account interface {
	// Create When creating a user, you need to create an Actor that supports the ActivityPub protocol at the same time.
	// First, need to verify the username and email address of the account are unique.
	Create() error

	Update() error

	// EditPassword Edit the account password, encrypt the password again, and exit all online clients.
	EditPassword() error

	// EditUsername Edit the account name will change both the
	// Account username field
	// and the Actor PreferredUsername field.
	EditUsername() error

	// Delete The account will be deleted by AccountID and the Actor data will be deleted at the same time.
	Delete() error

	GetAccountByID() (*Accounts, error)

	GetAccountByUsername() (*Accounts, error)
}
