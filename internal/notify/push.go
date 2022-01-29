package notify

type PushData struct {
	AccountID uint
	DeviceID  uint
	Data      []byte
}

func NewPush(accountID, deviceID uint, data []byte) *PushData {
	return &PushData{
		AccountID: accountID,
		DeviceID:  deviceID,
		Data:      data,
	}
}

func (p *PushData) Push() error {
	//d, err := device.NewDevicesByID(p.AccountID, p.DeviceID).GetDevice()
	//if err != nil {
	//	return err
	//}
	//// Get subscription by device ID.
	//n, err := NewNotifiesByDeviceID(d.ID).Get()
	//if err != nil {
	//	return err
	//}
	//if err := push.NewSubscription(p.DeviceID, n.Endpoint, n.Auth, n.P256dh, d.PublicKey, d.PrivateKey, p.Data).Send(); err != nil {
	//	return err
	//}
	return nil
}
