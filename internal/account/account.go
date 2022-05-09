package account

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Accounts is a struct for account.
type Accounts struct {

	// AccountID is a unique identifier for the account.
	gorm.Model

	// Username is the primaryKey of the database which is unique,
	// in the account system of this instance is must be added during
	// the creation process to ensure the correctness of the data.
	Username string `gorm:"primaryKey;type:text;username;unique" validate:"required,min=4,max=16"`

	// Mail When registering, the user is required to provide an email,
	// and an error needs to be returned when the email is not in the
	// correct format. It is also unique in the account system.
	Mail string `gorm:"index;type:text;mail;unique" validate:"required,email"`

	// Password must be encrypted and saved. The length of the password needs to be verified
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// ActorID is used for compatibility with the ActivityPub protocol
	// to connect to the actor table by ID.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// IsPrivate sets whether the account is private or not,
	// it is a social extension that is set by the user to make the
	// account public or not.
	IsPrivate bool `gorm:"type:boolean;is_private"`
}

type accounts interface {
	IsExist() bool
	Create(publicKey string) error
	DeleteAccount(password string) error
	EditUsername(username string) error
	EditPassword(newPassword string) error
	EditEmail(mail string) error
}

// NewUsername ...
func NewUsername(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

// IsExist ...
func (a *Accounts) IsExist() bool {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).Where("username = ? ", a.Username).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return ok
	}
	return false
}

// GetAccountByUsername ...
func (a *Accounts) GetAccountByUsername() (*Accounts, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(AccountsTable).
		Where("username = ?", a.Username).First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// NewAccounts ...
func NewAccounts(actorID uint, username, mail, password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: string(hash),
		ActorID:  actorID,
	}
}

// NewCreateAccounts ...
func NewCreateAccounts(username, mail, password string) *Accounts {
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: password,
	}
}

// Create Accounts...
func (a *Accounts) Create(publicKey string) error {
	if err := validator.New().Struct(a); err != nil {
		return errors.New("FAILED_TO_VALIDATOR")
	}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACTOR_DATABASE")
	}

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACCOUNT_DATABASE")
	}

	if err := db.Debug().Table(AccountsTable).
		Where("username = ? ", a.Username).Or("mail = ?", a.Mail).
		First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.New("THE_USERNAME_OR_MAIL_ALREADY_EXISTS")
		}
	}

	// Create actor.
	actor, err := NewActors(a.Username, publicKey, "Person").Create()
	if err != nil {
		return err
	}

	v := NewAccounts(actor.ID, a.Username, a.Mail, a.Password)
	if err := db.Debug().Table(AccountsTable).
		Create(&v).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}
	return nil
}

// NewDeleteAccount ...
func NewDeleteAccount(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

// DeleteAccount ...
func (a *Accounts) DeleteAccount(password string) error {
	// Verify account.
	v, err := NewVerify(a.Username).Verify(password)
	if err != nil {
		return err
	}

	db := cockroach.GetDB()

	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", v.ActorID).Unscoped().Delete(&Actors{}).Error; err != nil {
		return err
	}

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", v.ID).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return err
	}

	return nil
}

func NewEditAccountID(id uint) *Accounts {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}

// EditUsername ...
func (a *Accounts) EditUsername(username string) error {
	if ok := NewUsername(username).IsExist(); ok {
		return errors.New("FAILED_TO_VALIDATOR")
	}
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", a.ID).First(&a).Update("username", username).Error; err != nil {
		return err
	}

	address := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), username)

	inbox := fmt.Sprintf("%s/inbox", address)

	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", a.ActorID).
		Update("preferred_username", username).
		Update("inbox", inbox).
		Update("address", address).Error; err != nil {
		return err
	}

	return nil
}

func (a *Accounts) EditEmail(mail string) error {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", a.ID).
		Update("mail", mail).
		Error; err != nil {
		return err
	}

	return nil
}

func NewEditPassword(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

// EditPassword ...
func (a *Accounts) EditPassword(password, newPassword string) error {
	v, err := NewVerify(a.Username).Verify(password)
	if err != nil {
		return nil
	}

	db := cockroach.GetDB()
	hash, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", v.ID).Update("password", hash).Error; err != nil {
		return err
	}

	return nil
}
