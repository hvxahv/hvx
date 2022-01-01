package notify

import (
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/push"
)

type PushData struct {
	DeviceID uint
	Data     []byte
}

func NewPush(deviceID uint, data []byte) *PushData {
	return &PushData{
		DeviceID: deviceID,
		Data:     data,
	}
}

func (p *PushData) Push() error {
	d, err := accounts.NewDevicesByID(p.DeviceID).GetDevicesByID()
	if err != nil {
		return err
	}
	// Get subscription by device ID.
	n, err := NewNotifiesByDeviceID(d.DeviceID).Get()
	if err != nil {
		return err
	}
	if err := push.NewSubscription(p.DeviceID, n.Endpoint, n.Auth, n.P256dh, d.PublicKey, d.PrivateKey, p.Data).Send(); err != nil {
		return err
	}
	return nil
}
