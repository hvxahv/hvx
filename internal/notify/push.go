package notify

import "github.com/hvxahv/hvxahv/pkg/push"

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
	// Get subscription by device ID.
	n, err := NewNotifiesByDeviceID(p.DeviceID).Get()
	if err != nil {
		return err
	}
	if err := push.NewSubscription(p.DeviceID, n.Endpoint, n.Auth, n.P256dh, n.PublicKey, n.PrivateKey).Send(); err != nil {
		return err
	}
	return nil
}
