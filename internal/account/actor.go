package account

import (
	"fmt"

	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Actors struct {
	gorm.Model

	PreferredUsername string `gorm:"primaryKey;type:text;preferred_username;"`
	Domain            string `gorm:"index;type:text;domain"`
	Avatar            string `gorm:"type:text;avatar"`
	Name              string `gorm:"type:text;name"`
	Summary           string `gorm:"type:text;summary"`
	Inbox             string `gorm:"type:text;inbox"`
	Address           string `gorm:"index;test;address"`
	PublicKey         string `gorm:"type:text;public_key"`
	ActorType         string `gorm:"type:text;actor_type"`
	IsRemote          bool   `gorm:"type:boolean;is_remote"`
}

func (a *Actors) SetActorPreferredUsername(preferredUsername string) *Actors {
	a.PreferredUsername = preferredUsername
	return a
}

func (a *Actors) SetActorName(name string) *Actors {
	a.Name = name
	return a
}

func (a *Actors) SetActorAvatar(avatar string) *Actors {
	a.Avatar = avatar
	return a
}

func (a *Actors) SetActorSummary(summary string) *Actors {
	a.Summary = summary
	return a
}

func (a *Actors) Create() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACTOR_DATABASE")
	}

	if err := db.Debug().Table("actors").Create(&a).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	if err := db.Debug().Table("accounts").Where("username = ?", a.PreferredUsername).Update("actor_id", a.ID).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Actors) Update() error {
	if a.PreferredUsername != "" {
		return errors.New("PLEASE_USE_THE_SET_ACTOR_PREFERRED_USERNAME_METHOD_TO_UPDATE_THE_PREFERRED_USERNAME")
	}
	db := cockroach.GetDB()
	err := db.Debug().Table("actors").Where("id = ?", a.ID).Updates(&a).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Actors) EditActorPreferredUsername() error {
	db := cockroach.GetDB()
	err := db.Debug().Table("actors").Where("id = ?", a.ID).Update("preferred_username", a.PreferredUsername).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *Actors) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("id = ?", a.ID).Unscoped().Delete(&Actors{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Actors) GetActorsByPreferredUsername() (*[]Actors, error) {
	db := cockroach.GetDB()

	var ac []Actors
	if err := db.Debug().Table("actors").Where("preferred_username = ?", a.PreferredUsername).Find(&ac).Error; err != nil {
		return nil, err
	}

	return &ac, nil
}

func (a *Actors) GetActorByAccountUsername() (*Actors, error) {
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

func (a *Actors) GetActorByID() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Actors) GetActorByAddress() (*Actors, error) {
	db := cockroach.GetDB()

	err := db.Debug().Table("actors").Where("address = ?", a.Address).First(&a).Error
	if err != nil {
		ok := cockroach.IsNotFound(err)
		if ok {
			return nil, err
		}
	}
	return a, nil
}

func NewActorsAddress(address string) *Actors {
	return &Actors{Address: address}
}

func NewActorsAccountUsername(username string) *Actors {
	return &Actors{PreferredUsername: username}
}

func NewActorsPreferredUsername(preferredUsername string) *Actors {
	return &Actors{PreferredUsername: preferredUsername}
}

// NewAddActors Add an Actor from remote and set IsRemote to true.
func NewAddActors(preferredUsername, Domain, Avatar, Name, Summary, Inbox, address, PublicKey, ActorType string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            Domain,
		Avatar:            Avatar,
		Name:              Name,
		Summary:           Summary,
		Inbox:             Inbox,
		Address:           address,
		PublicKey:         PublicKey,
		ActorType:         ActorType,
		IsRemote:          true,
	}
}

func NewActorsID(id uint) *Actors {
	return &Actors{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewActors(preferredUsername, publicKey, actorType string) *Actors {
	domain := viper.GetString("localhost")

	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Inbox:             fmt.Sprintf("https://%s/u/%s/inbox", domain, preferredUsername),
		Address:           fmt.Sprintf("https://%s/u/%s", domain, preferredUsername),
		PublicKey:         publicKey,
		ActorType:         actorType,
		IsRemote:          false,
	}
}

type Actor interface {

	// Create actors field.
	Create() (*Actors, error)

	// Update Edit actors fields.
	Update() error

	EditActorPreferredUsername() error

	// Delete actor data by actor ID is usually called by account delete  method.
	Delete() error

	// GetActorsByPreferredUsername Get the Actor collection by PreferredUsername.
	GetActorsByPreferredUsername() (*[]Actors, error)

	// GetActorByAccountUsername Get unique Actor by username.
	GetActorByAccountUsername() (*Actors, error)

	GetActorByID() (*Actors, error)

	GetActorByAddress() (*Actors, error)
}
