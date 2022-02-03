package notify

import (
	pb "github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"golang.org/x/net/context"
	"strconv"
)

func (n *notify) Subscription(ctx context.Context, in *pb.NewNotifySubscription) (*pb.NotifySubscriptionReply, error) {
	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Notifies{}); err != nil {
		return nil, err
	}

	v := NewNotifies(uint(id), in.Endpoint, in.P256Dh, in.Auth)
	if err := db.Debug().Table("notifies").Create(&v).Error; err != nil {
		return nil, err
	}

	return &pb.NotifySubscriptionReply{Code: "200", Reply: "ok"}, nil
}

func NewNotifies(deviceID uint, endpoint, p256dh, auth string) *Notifies {
	return &Notifies{
		DeviceID: deviceID,
		Endpoint: endpoint,
		P256dh:   p256dh,
		Auth:     auth,
	}
}
