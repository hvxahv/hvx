package accounts

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Actors struct {
	gorm.Model

	PreferredUsername string `gorm:"primaryKey;type:text;preferredUsername;"`
	Domain            string `gorm:"index;type:text;domain"`
	Avatar            string `gorm:"type:text;avatar"`
	Name              string `gorm:"type:text;name"`
	Summary           string `gorm:"type:text;summary"`
	Inbox             string `gorm:"type:text;inbox"`
	Url               string `gorm:"index;test;url"`
	PublicKey         string `gorm:"type:text;public_key"`

	// ID returned after completing the registration of the matrix account.
	MatrixID    string `gorm:"type:text;matrix_id;unique"`
	MatrixToken string `gorm:"type:text;matrix_token"`

	// Whether it is a robot or other type of account
	ActorType string `gorm:"type:text;actor_type"`

	// Set whether it is a remote actor.
	IsRemote bool `gorm:"type:boolean;is_remote"`
}

func (a *Actors) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("id = ?", a.ID).Unscoped().Delete(&Actors{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Actors) GetActorByUri() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("url = ?", a.Url).First(&a).Error; err != nil {
		ok := cockroach.IsNotFound(err)
		if ok {
			return nil, err
		}
	}
	return a, nil
}

func NewActorUri(uri string) *Actors {
	return &Actors{Url: uri}
}

func (a *Actors) AddActor() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Create(&a).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	return a, nil
}

// NewAddActor Add an Actor from remote and set IsRemote to true.
func NewAddActor(PreferredUsername, Domain, Avatar, Name, Summary, Inbox, Url, PublicKey, MatrixID, ActorType string) *Actors {
	return &Actors{
		PreferredUsername: PreferredUsername,
		Domain:            Domain,
		Avatar:            Avatar,
		Name:              Name,
		Summary:           Summary,
		Inbox:             Inbox,
		Url:               Url,
		PublicKey:         PublicKey,
		MatrixID:          MatrixID,
		ActorType:         ActorType,
		IsRemote:          true,
	}
}

func (a *Actors) FindActorByAccountUsername() (*Actors, error) {
	db := cockroach.GetDB()

	acct, err := NewAccountsUsername(a.PreferredUsername).GetAccountByUsername()
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

func (a *Actors) GetByID() (*Actors, error) {
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

func NewActors(preferredUsername, password, publicKey, actorType string) *Actors {
	domain := viper.GetString("localhost")

	//id, err := matrix.NewAuth(preferredUsername, password).Register()
	//if err != nil {
	//	log.Println("FAILED TO REGISTER TO THE MATRIX ACCOUNT.")
	//}

	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Inbox:             fmt.Sprintf("https://%s/u/%s/inbox", domain, preferredUsername),
		Url:               fmt.Sprintf("https://%s/u/%s", domain, preferredUsername),
		PublicKey:         publicKey,
		MatrixID:          "",
		ActorType:         actorType,
		IsRemote:          false,
	}
}

func (a *Actors) NewActor() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_DATABASE")
	}

	if err := db.Debug().Table("actors").Create(&a).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	return a, nil
}

type Actor interface {
	// NewActor Create new actors data and add the returned ID to the accounts field.
	NewActor() (*Actors, error)

	AddActor() (*Actors, error)

	// FindByPreferredUsername Find the Actor collection by PreferredUsername.
	FindByPreferredUsername() (*[]Actors, error)

	FindActorByAccountUsername() (*Actors, error)
	GetByID() (*Actors, error)
	GetActorByUri() (*Actors, error)

	Update() error

	Delete() error
}
