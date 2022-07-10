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

	// ActorID is the unique identifier for the actor.
	gorm.Model

	// PreferredUsername is the preferred username for the actor.
	// The username is stored for multiple instances, so it is repeatable.
	PreferredUsername string `gorm:"primaryKey;type:text;preferred_username;"`

	// Domain name of the instance,
	// for example: https://example.com, example.com is the domain name.
	Domain string `gorm:"index;type:text;domain"`

	// Avatar is the URL of the avatar image.
	// It is not required.
	// It can also be an empty string.
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

	// IsRemote is a flag that indicates whether the actor is a remote actor.
	// Remote actors are still stored in the database.
	// They are used to indicating whether the actor is in the current instance or not.
	IsRemote bool `gorm:"type:boolean;is_remote"`
}

type Actor interface {
	Create() (*Actors, error)
	GetActorsByPreferredUsername() ([]*Actors, error)
	AddActor() error
	GetActorByUsername() (*Actors, error)
	Edit() error
	Delete() error
}

// NewActors creates a new instance of Actors.
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

func NewAddActors(preferredUsername, host, avatar, name, summary, inbox, address, publicKey, actorType string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
		Domain:            host,
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

func (a *Actors) AddActor() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(ActorsTable).
		Create(&a).Error; err != nil {
		return err
	}

	return nil
}

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
			actor := NewAddActors(
				f.PreferredUsername,
				h.Host,
				f.Icon.Url,
				f.Name,
				f.Summary,
				f.Inbox,
				a.Address,
				f.PublicKey.PublicKeyPem,
				f.Type,
			)
			if err := actor.AddActor(); err != nil {
				return nil, err
			}
			return actor, nil
		}
	}
	return a, nil
}

func NewAccountUsername(preferredUsername string) *Actors {
	return &Actors{
		PreferredUsername: preferredUsername,
	}
}

func (a *Actors) GetActorByUsername() (*Actors, error) {
	domain := viper.GetString("domain")
	db := cockroach.GetDB()
	var actor Actors
	if err := db.Debug().Table(ActorsTable).
		Where("preferred_username = ? AND domain = ?", a.PreferredUsername, domain).First(&actor).Error; err != nil {
		return nil, err
	}
	return &actor, nil
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

func NewActorId(actorId uint) *Actors {
	return &Actors{
		Model: gorm.Model{
			ID: actorId,
		},
	}
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
	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", a.ID).Delete(&a).Error; err != nil {
		return err
	}
	return nil

}
