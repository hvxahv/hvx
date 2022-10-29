package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hvxahv/hvx/activitypub"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/url"
)

// Actors are used to support the ActivityPub protocol,
// see the annotated links below to see more documentation on Actor.
// https://www.w3.org/TR/activitypub/#actor-objects
// Integrated in the account system, it can be used to represent account details, such as name, profile, etc.

// Actors is a struct that represents a role in the ActivityPub social system.
type Actors struct {
	gorm.Model

	// PreferredUsername is the preferred username for the actor.
	// The username is stored for multiple instances, so it is repeatable.
	PreferredUsername string `gorm:"primaryKey;type:text;preferred_username;"`

	// Domain name of the instance,
	// for example: https://example.com, example.com is the domain name.
	Domain string `gorm:"index;type:text;domain"`

	// Avatar is the URL of the avatar image. it can also be an empty string.
	Avatar string `gorm:"type:text;avatar"`

	// Name is the name of the actor, for example: John Doe.
	Name string `gorm:"type:text;name"`

	// Summary is a description of the actor. It will become a comprehensive description of the actor.
	Summary string `gorm:"type:text;summary"`

	// Inbox is the URL of the actor's inbox,
	// for example: https://halfmemories.com/u/hvturingga/inbox.
	Inbox string `gorm:"type:text;inbox"`

	// Outbox is the URL of the actor's outbox,
	// for example: https://halfmemories.com/u/hvturingga/outbox.
	Address string `gorm:"index;test;address"`

	// PublicKey is the public key of the actor. It is used to verify the signature of the actor.
	PublicKey string `gorm:"type:text;public_key"`

	// ActorType is the type of the actor. It is used to determine the type of the actor.
	ActorType string `gorm:"type:text;actor_type"`

	// IsRemote is a flag indicating whether the actor is a remote actor or not.
	// Remote participants are still stored in the database.
	// Used to indicate whether the actor's account is in the current instance.
	IsRemote bool `gorm:"type:boolean;is_remote"`
}

type Actor interface {
	// IsExist Actor exists or not.
	IsExist() (*Actors, bool)

	// Create Actor.
	Create() (*Actors, error)

	// Get Actor.
	Get() (*Actors, error)

	// GetActorsByPreferredUsername Get Actors by preferred username.
	GetActorsByPreferredUsername() ([]*Actors, error)

	// AddActor Add Actor.
	// Used when saving users from other instances to this instance.
	AddActor() (*Actors, error)

	// GetActorByUsername Get Actor by username.
	GetActorByUsername() (*Actors, error)

	// Edit Actor.
	Edit() error

	// Delete Actor.
	Delete() error
}

// NewActorsIsExist Determines if the Actor exists in the constructor of the current instance.
func NewActorsIsExist(domain, preferredUsername string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
	}
}

func (a *Actors) IsExist() (*Actors, bool) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ActorsTable).
		Where("preferred_username = ? AND domain = ? ", a.PreferredUsername, a.Domain).
		First(&a); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return a, !ok
	}
	return nil, false
}

// NewActors The constructor for creating a new Actor.
func NewActors(preferredUsername, publicKey, actorType string) *Actors {
	domain := viper.GetString("domain")
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

func NewChannels(preferredUsername, publicKey, actorType string) *Actors {
	domain := viper.GetString("domain")
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Inbox:             fmt.Sprintf("https://%s/c/%s/inbox", domain, preferredUsername),
		Address:           fmt.Sprintf("https://%s/c/%s", domain, preferredUsername),
		PublicKey:         publicKey,
		ActorType:         actorType,
		IsRemote:          false,
	}
}

func (a *Actors) Create() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACTOR_DATABASE")
	}

	if err := db.Debug().Table(ActorsTable).
		Create(&a).
		Error; err != nil {
		return nil, err
	}
	return &Actors{
		Model: gorm.Model{
			ID: a.ID,
		},
	}, nil
}

// NewActorsId Instantiates the constructor of an ActorsId.
func NewActorsId(id uint) *Actors {
	return &Actors{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (a *Actors) Get() (*Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// NewPreferredUsername The constructor creates a new PreferredUsername.
func NewPreferredUsername(preferredUsername string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
	}
}

func (a *Actors) GetActorsByPreferredUsername() ([]*Actors, error) {
	db := cockroach.GetDB()

	var actors []*Actors
	if err := db.Debug().Table(ActorsTable).
		Where("preferred_username = ?", a.PreferredUsername).Find(&actors).Error; err != nil {
		return nil, err
	}
	return actors, nil
}

// NewAddActors The constructor is used to add a new Actor.
func NewAddActors(preferredUsername, domain, avatar, name, summary, inbox, address, publicKey, actorType string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
		Avatar:            avatar,
		Name:              name,
		Summary:           summary,
		Inbox:             inbox,
		Address:           address,
		PublicKey:         publicKey,
		ActorType:         actorType,
		IsRemote:          true,
	}
}

func (a *Actors) AddActor() (*Actors, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(ActorsTable).
		Create(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

// NewActorAddress The constructor is used to get an Actor by address.
func NewActorAddress(address string) *Actors {
	return &Actors{
		Address: address,
	}
}

func (a *Actors) GetActorByAddress() (*Actors, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("actors").Where("address = ?", a.Address).First(&a).Error; err != nil {
		if cockroach.IsNotFound(err) {
			res, err := resty.New().R().
				SetHeader("Content-Type", "application/activitypub+json; charset=utf-8").
				SetHeader("Accept", "application/ld+json").
				EnableTrace().
				Get(a.Address)
			if err != nil {
				return nil, err
			}

			var f *activitypub.Actor

			if err = json.Unmarshal(res.Body(), &f); err != nil {
				return nil, err
			}

			h, err := url.Parse(a.Address)
			if err != nil {
				return nil, err
			}

			actor, err := NewAddActors(
				f.PreferredUsername,
				h.Host,
				f.Icon.Url,
				f.Name,
				f.Summary,
				f.Inbox,
				a.Address,
				f.PublicKey.PublicKeyPem,
				f.Type,
			).AddActor()
			if err != nil {
				return nil, err
			}
			return actor, nil
		}
	}
	return a, nil
}

// NewAccountUsername The constructor is used to get an Actor by username.
func NewAccountUsername(preferredUsername string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            viper.GetString("domain"),
	}
}

func NewPreferredUsernameAndDomain(preferredUsername, domain string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            domain,
	}
}

func (a *Actors) GetActorByUsername() (*Actors, error) {
	db := cockroach.GetDB()
	var actor Actors
	if err := db.Debug().Table(ActorsTable).
		Where("preferred_username = ? AND domain = ?", a.PreferredUsername, a.Domain).First(&actor).Error; err != nil {
		return nil, err
	}
	return &actor, nil
}

func (a *Actors) Edit() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", a.ID).Updates(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Actors) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ActorsTable).
		Where("id = ?", a.ID).
		Unscoped().
		Delete(&Actors{}).Error; err != nil {
		return err
	}
	return nil

}
