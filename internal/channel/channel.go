package channel

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hvxahv/hvxahv/api/account/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/identity"
	"gorm.io/gorm"
)

const (
	// ErrNotFound is returned when a channel is not found.
	ErrNotFound   = "CHANNEL_NOT_FOUND"
	ChannelsTable = "channels"
)

type Channels struct {
	gorm.Model

	// ActorID This ID is associated with the Actor table, as an
	// ActivityPub Actor service. This ID can be used
	// to retrieve data from the Actor table. as the channel information.
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
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Channels{}); err != nil {
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

	ch := NewChannels(uint(id), uint(aid), string(privateKey))
	if err := db.Debug().
		Table(ChannelsTable).
		Create(ch).
		Error; err != nil {
		return nil, err

	}
	administrator, err := c.AddAdministrator(ctx, &pb.AddAdministratorRequest{
		ChannelId:      strconv.Itoa(int(ch.ID)),
		AdminAccountId: in.AccountId,
		AddAdminId:     in.AccountId,
		IsOwner:        true,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateChannelResponse{Code: "200", Reply: administrator.Reply}, nil
}

func (c *channel) GetChannelsByAccountID(ctx context.Context, in *pb.GetChannelsByAccountIDRequest) (*pb.GetChannelsByAccountIDResponse, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Channels{}); err != nil {
		return nil, err
	}

	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	var channels []Channels
	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ?", uint(aid)).
		Find(&channels).
		Error; err != nil {
		return nil, err
	}

	var chs []*pb.Channel
	for _, ch := range channels {
		chs = append(chs, &pb.Channel{
			Id:        strconv.Itoa(int(ch.ID)),
			ChannelId: strconv.Itoa(int(ch.ActorID)),
		})
	}

	return &pb.GetChannelsByAccountIDResponse{Code: "200", Channels: chs}, nil
}

func (c *channel) DeleteChannel(ctx context.Context, in *pb.DeleteChannelRequest) (*pb.DeleteChannelResponse, error) {
	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.ChannelId,
		AccountId: in.AccountId,
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, fmt.Errorf("%s", NotAdmin)
	}
	client, err := account.GetActorClient()
	if err != nil {
		return nil, err
	}
	d, err := client.DeleteActor(ctx, &v1alpha1.DeleteActorRequest{
		AccountId: in.AccountId,
	})
	if err != nil && d.Code != "200" {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Channels{}); err != nil {
		return nil, err
	}

	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ? AND id = ?", uint(aid), uint(id)).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table(AdministrateTable).
		Where("admin_id = ? AND c_id = ? AND is_owner = ?", uint(aid), uint(id), true).
		Unscoped().
		Delete(&Administrates{}).
		Error; err != nil {
		return nil, err

	}
	return &pb.DeleteChannelResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) DeleteAllChannelsByAccountID(ctx context.Context, in *pb.DeleteAllChannelsByAccountIDRequest) (*pb.DeleteAllChannelsByAccountIDResponse, error) {
	db := cockroach.GetDB()

	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ?", uint(aid)).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return nil, err
	}
	return &pb.DeleteAllChannelsByAccountIDResponse{Code: "200", Reply: "ok"}, nil
}
