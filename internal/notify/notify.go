package notify

import (
	"context"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	pb "github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"strconv"
)

// Submit the client-registered data to hvxahv, and the pushed endpoint,
// p256dh and auth. hvxahv will send user notification message through the address.
// https://datatracker.ietf.org/doc/html/rfc6108

// https://developer.mozilla.org/en-US/docs/Web/API/Push_API
// {
//   "subscription":{
//     "endpoint":"https://sg2p.notify.windows.com/w/?token=BQYAAACLy9DPk5%2bkrOfxjm4cTtROW5peMRUg1m609l8lvgOje%2bzvQXtv2zW9pg286gAo8X67OlhkvCN7MC90MR6YkdBLk4aXHmiXd941QtkMLoDIv8ep2kXLgfKs7VEhJE1%2bXA9XOeoED%2brmFlojNhkxO%2b9N93cjob54Jo0nzfXUDUyOSUnXkgfcbnx7a0M4u9dExBCXIpfvIaUIvBmhqGhWgyW0KrgUv%2bj7R73SErOWdoYfOOISybv7Io55NAbrjRFRMIFzkepB3LLd2F2KZKQTg3o2f4nghRcM0qRqkdVRregsgl6eAFyzndhxVTff%2fBSiTP0%3d",
//     "expirationTime":null,
//     "keys":{
//       "p256dh":"BM9H6kiGNliWfhI23CrawUefVwsCYkIFZCsggtNTYSNy4Y5BzEJVrK3iM_0ZMzRndKYj2z7fXmBzxoQSrXnvxsQ",
//       "auth":"EJOFwKDlaVpDUV7uxcNPwg"
//     }
//   }
// }

type Notifies struct {
	gorm.Model

	DeviceID uint   `gorm:"primaryKey;type:bigint;device_id"`
	Endpoint string `gorm:"type:text;endpoint"`
	P256dh   string `gorm:"type:text;p256dh"`
	Auth     string `gorm:"type:text;auth"`
}

func (n *notify) Subscription(ctx context.Context, in *pb.NewNotifySubscription) (*pb.NotifySubscriptionReply, error) {
	client, err := device.NewDeviceClient()
	if err != nil {
		return nil, err
	}

	d := &v1alpha1.NewDeviceHash{Hash: in.DeviceHash}
	h, err := client.GetDeviceByHash(ctx, d)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(h.Id)
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

func NewNotifiesByDeviceID(id uint) *Notifies {
	return &Notifies{DeviceID: id}
}

func (n *Notifies) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Notifies{}); err != nil {
		return err
	}
	if err := db.Debug().Table("notifies").Create(&n).Error; err != nil {
		return err
	}
	return nil
}

func (n *Notifies) Get() (*Notifies, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("notifies").Where("device_id = ?", n.DeviceID).First(&n).Error; err != nil {
		return nil, err
	}
	return n, nil
}

type Notify interface {
	Create() error
	Get() (*Notifies, error)
}
