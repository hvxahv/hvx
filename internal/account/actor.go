package account

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-resty/resty/v2"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/gorm"
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

func (a *account) CreateActor(ctx context.Context, in *pb.CreateActorRequest) (*pb.CreateActorResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table("actors").
		Where("preferred_username = ? AND is_remote = ?", in.PreferredUsername, false).
		First(&a.Actors); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return nil, fmt.Errorf("ACTOR_ALREADY_EXISTS")
		}
	}
	actors := NewActors(in.PreferredUsername, in.PublicKey, in.ActorType)
	if err := db.Debug().
		Table("actors").
		Create(&actors).
		Error; err != nil {
		return nil, err
	}
	return &pb.CreateActorResponse{Code: "200", ActorId: strconv.Itoa(int(actors.ID))}, nil
}

func (a *account) GetActorByAccountUsername(ctx context.Context, in *pb.GetActorByAccountUsernameRequest) (*pb.AccountDataResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ? ", in.Username).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).First(&a.Actors).Error; err != nil {
		return nil, err
	}
	return &pb.AccountDataResponse{
		Id:                strconv.Itoa(int(a.Actors.ID)),
		PreferredUsername: a.Actors.PreferredUsername,
		Domain:            a.Actors.Domain,
		Avatar:            a.Actors.Avatar,
		Name:              a.Actors.Name,
		Summary:           a.Actors.Summary,
		Inbox:             a.Actors.Inbox,
		Address:           a.Actors.Address,
		PublicKey:         a.Actors.PublicKey,
		ActorType:         a.Actors.ActorType,
		IsRemote:          strconv.FormatBool(a.Actors.IsRemote),
	}, nil
}

func (a *account) GetActorsByPreferredUsername(ctx context.Context, in *pb.GetActorsByPreferredUsernameRequest) (*pb.GetActorsByPreferredUsernameResponse, error) {
	db := cockroach.GetDB()

	var actors []*pb.AccountDataResponse
	if err := db.Debug().Table("actors").Where("preferred_username = ?", in.PreferredUsername).Find(&actors).Error; err != nil {
		return nil, err
	}

	return &pb.GetActorsByPreferredUsernameResponse{Code: "200", Actors: actors}, nil
}

func (a *account) GetActorByAddress(ctx context.Context, in *pb.GetActorByAddressRequest) (*pb.AccountDataResponse, error) {
	db := cockroach.GetDB()
	actor := &Actors{}
	if err := db.Debug().Table("actors").Where("address = ?", in.Address).First(&actor).Error; err != nil {
		if cockroach.IsNotFound(err) {
			resp, err := resty.New().R().
				SetHeader("Content-Type", "application/activitypub+json; charset=utf-8").
				SetHeader("Accept", "application/ld+json").
				EnableTrace().
				Get(in.Address)
			if err != nil {
				return nil, err
			}

			fmt.Println(string(resp.Body()))
			var f *activitypub.Actor

			if err = json.Unmarshal(resp.Body(), &f); err != nil {
				return nil, err
			}

			h, err := url.Parse(in.Address)
			if err != nil {
				return nil, err
			}
			x := NewActorsAdd(f.PreferredUsername, h.Host, f.Icon.Url, f.Name, f.Summary, f.Inbox, in.Address, f.PublicKey.PublicKeyPem, f.Type)
			if err := db.Debug().Table("actors").Create(&x).Error; err != nil {
				return nil, err
			}
			return &pb.AccountDataResponse{
				Id:                strconv.Itoa(int(x.ID)),
				PreferredUsername: x.PreferredUsername,
				Domain:            x.Domain,
				Avatar:            x.Avatar,
				Name:              x.Name,
				Summary:           x.Summary,
				Inbox:             x.Inbox,
				Address:           x.Address,
				PublicKey:         x.PublicKey,
				ActorType:         x.ActorType,
				IsRemote:          strconv.FormatBool(x.IsRemote),
			}, nil
		}
		return nil, err
	}

	return &pb.AccountDataResponse{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

func (a *account) EditActor(ctx context.Context, in *pb.EditActorRequest) (*pb.EditActorResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ? ", in.AccountUsername).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	actor := new(Actors)
	if in.Avatar != "" {
		actor.SetActorAvatar(in.Avatar)
	}
	if in.Name != "" {
		actor.SetActorName(in.Name)
	}
	if in.Summary != "" {
		actor.SetActorSummary(in.Summary)
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).Updates(&actor).Error; err != nil {
		return nil, err
	}
	return &pb.EditActorResponse{Code: "200", Reply: "ok"}, nil
}

func NewActorsAdd(preferredUsername, host, avatar, name, summary, inbox, address, publicKey, actorType string) *Actors {
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

// NewActors creates a new instance of Actors.
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
