package accounts

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/matrix"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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

func NewAccounts(username string, mail string, password string) *Accounts {
	hash := security.GenPassword(password)
	return &Accounts{Username: username, Mail: mail, Password: hash}
}

func (a *Accounts) New() error {
	if err := validator.New().Struct(*a); err != nil {

	}

	// Before creating, first, check whether the user exists. If it does not exist, create the user.
	// If the account is found, it returns the user has created.
	// It will not be judged so detailed in the database. It only returns the information created by the user.
	// If more processing is required, detailed operations need to be done in the cache.

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("accounts").Where("username = ? ", a.Username).Or("mail = ?", a.Mail).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.New("THE_ACCOUNT_ALREADY_EXISTS")
		}
	}

	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("failed to generate public and private keys: %v", err)
	}

	id, err := NewActors(a.Username, a.Password, publicKey).NewActor()
	if err != nil {

	}

	a.ActorID = id
	a.PrivateKey = privateKey

	if err := db.Debug().Table("accounts").Create(&a).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	if err := db.Debug().Table("accounts").Create(&a).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	return nil
}

type Actors struct {
	gorm.Model
	PreferredUsername string `gorm:"type:text;preferredUsername;"`
	Domain            string `gorm:"type:text;domain"`

	Avatar    string `gorm:"type:text;avatar"`
	Name      string `gorm:"type:text;name"`
	Summary   string `gorm:"type:text;summary"`
	Inbox     string `gorm:"type:text;inbox"`
	PublicKey string `gorm:"type:text;public_key"`

	// ID returned after completing the registration of the matrix account.
	MatrixID string `gorm:"type:text;matrix_id;unique"`
	// Whether it is a robot or other type of account
	ActorType string `gorm:"type:text;actor_type"`
}

func NewActors(preferredUsername, password, publicKey string) *Actors {
	domain := viper.GetString("localhost")

	id, err := matrix.NewAuth(preferredUsername, password).Register()
	if err != nil {
		log.Println("Failed to register to the matrix account")
	}

	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Inbox:             fmt.Sprintf("https://%s/u/%s/inbox", domain, preferredUsername),
		PublicKey:         publicKey,
		MatrixID:          id,
	}
}

func (a Actors) NewActor() (uint, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return 0, errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("actors").Create(&a).Error; err != nil {
		return 0, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	return a.ID, nil
}

type Actor interface {
	NewActor() (uint, error)
}

type Account interface {
	New() error
}
