package channel

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Subscribes struct {
	gorm.Model

	ChannelID uint `gorm:"primaryKey;channel_id"`
	AccountID uint `gorm:"primaryKey;account_id"`
}

func isSubExist(db *gorm.DB, channelID, accountID uint) bool {
	var count int64
	db.Model(&Subscribes{}).Where("channel_id = ? AND account_id = ?", channelID, accountID).Count(&count)
	return count > 0
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
	if isSubExist(db, uint(cid), uint(aid)) {
		return nil, errors.New("SUBSCRIBER_ALREADY_EXIST")
	}
	if err := db.Debug().Table("subscribes").Create(&Subscribes{
		ChannelID: uint(cid),
		AccountID: uint(aid),
	}).Error; err != nil {
		return nil, err
	}
	return &pb.AddSubscriberResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) Unsubscribe(ctx context.Context, in *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	db := cockroach.GetDB()

	cid, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table("subscribes").
		Where("channel_id = ? AND account_id = ?", cid, aid).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return nil, err
	}
	return &pb.UnsubscribeResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) RemoveSubscriber(ctx context.Context, in *pb.RemoveSubscriberRequest) (*pb.RemoveSubscriberResponse, error) {
	db := cockroach.GetDB()

	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.ChannelId,
		AccountId: in.AdminId,
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, errors.New("NOT_CHANNEL_ADMINISTRATOR")
	}

	cid, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	sid, err := strconv.Atoi(in.SubscriberId)
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table("subscribes").
		Where("channel_id = ? AND account_id = ?", cid, sid).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return nil, err
	}
	return &pb.RemoveSubscriberResponse{Code: "200", Reply: "ok"}, nil
}

func (c *channel) GetAllSubscribers(ctx context.Context, in *pb.GetAllSubscribersRequest) (*pb.GetAllSubscribersResponse, error) {
	administrator, err := c.IsChannelAdministrator(ctx, &pb.IsChannelAdministratorRequest{
		ChannelId: in.ChannelId,
		AccountId: in.AdminId,
	})
	if err != nil {
		return nil, err
	}
	if !administrator.IsAdministrator {
		return nil, errors.New(NotAdmin)
	}
	db := cockroach.GetDB()

	cid, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	var subscribers []Subscribes
	if err := db.Debug().
		Table("subscribes").
		Where("channel_id = ?", cid).
		Find(&subscribers).
		Error; err != nil {
		return nil, err
	}
	var subscriberIds []string
	for _, i := range subscribers {
		subscriberIds = append(subscriberIds, strconv.Itoa(int(i.AccountID)))
	}
	return &pb.GetAllSubscribersResponse{Code: "200", Subscriber: subscriberIds}, nil
}
