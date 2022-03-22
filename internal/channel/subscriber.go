package channel

import (
	"context"
	"strconv"

	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Subscribes struct {
	gorm.Model

	ChannelID uint `gorm:"primaryKey;channel_id"`
	AccountID uint `gorm:"primaryKey;account_id"`
}

func (c *channel) AddSubscriber(ctx context.Context, in *pb.AddSubscriberRequest) (*pb.AddSubscriberResponse, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Subscribes{}); err != nil {
		return nil, err
	}
	cid, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	if err := db.Debug().Table("subscribes").Create(&Subscribes{
		ChannelID: uint(cid),
		AccountID: uint(aid),
	}).Error; err != nil {
		return nil, err
	}
	return &pb.AddSubscriberResponse{Code: "200", Reply: "ok"}, nil
}
