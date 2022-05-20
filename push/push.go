/**
                 Generic Event Delivery Using HTTP Push
                     draft-ietf-webpush-protocol-12

	Abstract

	A simple protocol for the delivery of real-time events to user agents
	is described.  This scheme uses HTTP/2 server push.
	https://datatracker.ietf.org/doc/html/draft-ietf-webpush-protocol-12
*/

package push

import (
	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	DeviceID string
	Endpoint string
	Auth     string
	P256dh   string

	// VAPID KEY
	PublicKey  string
	PrivateKey string

	Data []byte
}

type Data struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Icon  string `json:"icon"`
	Tag   string `json:"tag"`
}

func NewData(title string, body string, icon string, tag string) *Data {
	return &Data{Title: title, Body: body, Icon: icon, Tag: tag}
}

func NewSubscription(deviceID, endpoint, auth, p256dh, publicKey, privateKey string, data []byte) *Subscription {
	return &Subscription{
		DeviceID:   deviceID,
		Endpoint:   endpoint,
		Auth:       auth,
		P256dh:     p256dh,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Data:       data,
	}
}

func (s *Subscription) Send() error {
	sub := &webpush.Subscription{
		Endpoint: s.Endpoint,
		Keys: webpush.Keys{
			Auth:   s.Auth,
			P256dh: s.P256dh,
		},
	}

	// Send Notification.
	resp, err := webpush.SendNotification(s.Data, sub, &webpush.Options{
		Subscriber:      s.DeviceID, // Do not include "mailto:"
		VAPIDPublicKey:  s.PublicKey,
		VAPIDPrivateKey: s.PrivateKey,
		TTL:             30,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
