package accounts

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/matrix"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type Actors struct {
	gorm.Model

	PreferredUsername string `gorm:"primaryKey;type:text;preferredUsername;"`
	Domain            string `gorm:"type:text;domain"`
	Avatar            string `gorm:"type:text;avatar"`
	Name              string `gorm:"type:text;name"`
	Summary           string `gorm:"type:text;summary"`
	Inbox             string `gorm:"type:text;inbox"`
	PublicKey         string `gorm:"type:text;public_key"`

	// ID returned after completing the registration of the matrix account.
	MatrixID    string `gorm:"type:text;matrix_id;unique"`
	MatrixToken string `gorm:"type:text;matrix_token"`

	// Whether it is a robot or other type of account
	ActorType string `gorm:"type:text;actor_type"`
}

func (a *Actors) FindActorByAccountUsername() (*Actors, error) {
	db := cockroach.GetDB()

	acct, err := NewAccountsName(a.PreferredUsername).FindAccountByUsername()
	if err != nil {
		return nil, err
	}

	if err := db.Debug().Table("actors").Where("id = ?", acct.ActorID).First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Actors) Update() error {
	db := cockroach.GetDB()

	err := db.Debug().Table("actors").Where("id = ?", a.ID).Updates(&a).Error
	if err != nil {
		return err
	}

	return nil
}

func NewActorID(id uint) *Actors {
	return &Actors{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (a *Actors) FindActorByID() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func NewActorsPreferredUsername(preferredUsername string) *Actors {
	return &Actors{PreferredUsername: preferredUsername}
}

func (a *Actors) FindByPreferredUsername() (*[]Actors, error) {
	db := cockroach.GetDB()

	var ac []Actors
	if err := db.Debug().Table("actors").Where("preferred_username = ?", a.PreferredUsername).Find(&ac).Error; err != nil {
		return nil, err
	}

	return &ac, nil
}

func NewActors(preferredUsername, password, publicKey string) *Actors {
	domain := viper.GetString("localhost")

	id, err := matrix.NewAuth(preferredUsername, password).Register()
	if err != nil {
		log.Println("FAILED TO REGISTER TO THE MATRIX ACCOUNT.")
	}

	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Inbox:             fmt.Sprintf("https://%s/u/%s/inbox", domain, preferredUsername),
		PublicKey:         publicKey,
		MatrixID:          id,
	}
}

func (a *Actors) NewActor() (uint, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return 0, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_DATABASE")
	}

	if err := db.Debug().Table("actors").Create(&a).Error; err != nil {
		return 0, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	return a.ID, nil
}

type Actor interface {
	// NewActor Create new actors data and add the returned ID to the accounts field.
	NewActor() (uint, error)

	// FindByPreferredUsername Find the Actor collection by PreferredUsername.
	FindByPreferredUsername() (*[]Actors, error)

	FindActorByAccountUsername() (*Actors, error)
	FindActorByID() (*Actors, error)

	Update() error
}
