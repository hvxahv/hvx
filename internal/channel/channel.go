package channel

import (
	"context"
	"strconv"

	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/identity"
	"gorm.io/gorm"
)

type Channels struct {
	gorm.Model

	ActorID    uint   `gorm:"primaryKey;type:bigint;actor_id"`
	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id"`
	PrivateKey string `gorm:"private_key;type:text;private_key"`
}

func NewChannels(actorID uint, accountID uint, privateKey string) *Channels {
	return &Channels{ActorID: actorID, AccountID: accountID, PrivateKey: privateKey}
}

func (c *channel) CreateChannel(ctx context.Context, in *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	client, err := account.GetActorClient()
	if err != nil {
		return nil, err
	}

	privateKey, publicKey := identity.GenRsaKey()
	actor, err := client.CreateActor(context.Background(), &v1alpha1.CreateActorRequest{
		PreferredUsername: in.PreferredUsername,
		PublicKey:         string(publicKey),
		ActorType:         "Service",
	})
	if err != nil {
		return nil, err
	}

	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(actor.ActorId)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Channels{}); err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table("channels").
		Create(NewChannels(uint(id), uint(aid), string(privateKey))).
		Error; err != nil {
		return nil, err

	}
	return &pb.CreateChannelResponse{Code: "200", Reply: "ok"}, nil
}
