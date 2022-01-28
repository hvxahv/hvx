package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"strconv"

	"github.com/hvxahv/hvxahv/pkg/cockroach"
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

	return &pb.ActorsData{Code: "200", Actors: actors}, nil
}

func (a *account) AddActor(ctx context.Context, in *pb.ActorData) (*pb.Reply, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Create(&Actors{
		PreferredUsername: in.PreferredUsername,
		Domain:            in.Domain,
		Avatar:            in.Avatar,
		Name:              in.Name,
		Summary:           in.Summary,
		Inbox:             in.Inbox,
		Address:           in.Address,
		PublicKey:         in.PublicKey,
		ActorType:         in.ActorType,
		IsRemote:          true,
	}).Error; err != nil {
		return nil, err
	}
	return &pb.Reply{Code: "200", Reply: "ok"}, nil
}

func (a *account) EditActor(ctx context.Context, in *pb.NewEditActor) (*pb.Reply, error) {

	d := &pb.NewAccountUsername{Username: in.AccountUsername}
	acct, err := a.GetAccountByUsername(context.Background(), d)
	if err != nil {
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

	db := cockroach.GetDB()

	if err := db.Debug().Table("actors").Where("id = ?", acct.ActorId).Updates(&actor).Error; err != nil {
		return nil, err
	}
	return &pb.Reply{Code: "200", Reply: "ok"}, nil
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
