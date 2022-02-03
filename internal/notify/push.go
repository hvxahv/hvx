package notify

import (
	"context"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/push"
	"github.com/pkg/errors"
	"strconv"
)

func (n *notify) Push(ctx context.Context, in *pb.NewNotifyPush) (*pb.NotifyPushReply, error) {
	client, err := device.NewDeviceClient()
	if err != nil {
		return nil, err
	}

	de, err := client.GetDeviceByID(ctx, &v1alpha1.NewDeviceID{Id: in.DeviceId})
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}
	if err := db.Debug().Table("notifies").Where("device_id = ?", uint(id)).First(&n.Notifies).Error; err != nil {
		return nil, err
	}

	if err := push.NewSubscription(in.DeviceId, n.Notifies.Endpoint, n.Notifies.Auth, n.Notifies.P256dh, de.PublicKey, de.PrivateKey, in.Data).Send(); err != nil {
		return nil, errors.Errorf("PUSH_ERROR_%s ", err)
	}
	return &pb.NotifyPushReply{Code: "200", Reply: "ok"}, nil
}
