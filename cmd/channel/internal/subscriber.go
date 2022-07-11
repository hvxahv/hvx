package internal

import (
	"context"
	"strconv"

	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"

	"gorm.io/gorm"
)

const (
	SubscribesTable = "subscribes"
)

type Subscribes struct {
	gorm.Model

	ChannelId uint `gorm:"primaryKey;channel_id"`
	ActorId   uint `gorm:"primaryKey;actor_id"`
}

type Subscribe interface {
	IsSubscriber() bool
	AddSubscriber() error
}

func NewSubscribe(channelId, actorId uint) *Subscribes {
	return &Subscribes{
		ChannelId: channelId,
		ActorId:   actorId,
	}
}

func (sub *Subscribes) IsSubscriber() bool {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND actor_id = ?", sub.ChannelId, sub.ActorId).
		First(&Subscribes{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return true
		}
	}

	return false
}

func (sub *Subscribes) AddSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	is := sub.IsSubscriber()
	if is {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Subscribes{}); err != nil {
		return errors.NewDatabaseCreate("subscribes")
	}

	if err := db.Debug().
		Table(SubscribesTable).
		Create(&sub).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) RemoveSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	is := sub.IsSubscriber()
	if is {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND actor_id = ?", sub.ChannelId, sub.ActorId).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) Unsubscribe() error {

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
		AdminId:   in.AdminId,
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
		AdminId:   in.AdminId,
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
