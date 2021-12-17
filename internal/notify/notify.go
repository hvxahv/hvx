package notify

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

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

	DeviceID   uint   `gorm:"primaryKey;type:bigint;device_id"`
	Endpoint   string `gorm:"type:text;endpoint"`
	P256dh     string `gorm:"type:text;p256dh"`
	Auth       string `gorm:"type:text;auth"`
	PublicKey  string `gorm:"type:text;public_key"`
	PrivateKey string `gorm:"type:text;private_key"`
}

func NewNotifies(deviceID uint, endpoint, p256dh, auth, public_key, private_key string) *Notifies {
	return &Notifies{
		DeviceID:   deviceID,
		Endpoint:   endpoint,
		P256dh:     p256dh,
		Auth:       auth,
		PublicKey:  public_key,
		PrivateKey: private_key,
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
