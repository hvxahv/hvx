package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"golang.org/x/net/context"
	"strconv"

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

/****


grpc
*/

func (a *account) GetActorByAccountUsername(ctx context.Context, in *pb.NewAccountUsername) (*pb.ActorData, error) {
	d := &pb.NewAccountUsername{Username: in.Username}
	acct, err := a.GetAccountByUsername(context.Background(), d)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()

	id, err := strconv.Atoi(acct.ActorId)
	if err != nil {
		return nil, err
	}
	if err := db.Debug().Table("actors").Where("id = ?", uint(id)).First(&a.Actors).Error; err != nil {
		return nil, err
	}

	return &pb.ActorData{
		Id:                acct.AccountId,
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

func (a *account) GetActorsByPreferredUsername(ctx context.Context, in *pb.NewActorPreferredUsername) (*pb.ActorsData, error) {
	db := cockroach.GetDB()

	var actors []*pb.ActorData
	if err := db.Debug().Table("actors").Where("preferred_username = ?", in.PreferredUsername).Find(&actors).Error; err != nil {
		return nil, err
	}

	//for _, i := range actors {
	//	from, err := account.NewActorID(i.FromID).GetByActorID()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	inboxes := Inboxes{
	//		ID:           strconv.FormatUint(uint64(i.ID), 10),
	//		From:         from.Name,
	//		FromID:       strconv.FormatUint(uint64(i.FromID), 10),
	//		ActivityType: i.ActivityType,
	//		ActivityID:   i.ActivityID,
	//	}
	//	ibx = append(ibx, inboxes)
	//}

	fmt.Println(actors)

	return &pb.ActorsData{Code: "200", Actors: actors}, nil
}
