package internal

import (
	"encoding/json"
	"github.com/hvxahv/hvx/cache"
	"strconv"
)

type Dh struct {
	DeviceId       uint
	ObjectDeviceId uint
	IV             string
	PublicKey      string
	PrivateKey     string
}

func NewDh(deviceId uint, objectDeviceId uint, IV string, publicKey string) *Dh {
	return &Dh{DeviceId: deviceId, ObjectDeviceId: objectDeviceId, IV: IV, PublicKey: publicKey}
}

type DH interface {
	GetPrivateKey() error
	GetDH() (*Dh, error)
	SendPrivateKey() error
}

func (dh *Dh) GetPrivateKey() error {
	marshal, err := json.Marshal(dh)
	if err != nil {
		return err
	}

	if err := cache.NewCache(1).SETDH(strconv.Itoa(int(dh.ObjectDeviceId)), marshal); err != nil {
		return err
	}
	return nil
}

func NewDHDeviceId(deviceId uint) *Dh {
	return &Dh{DeviceId: deviceId}
}

func (dh *Dh) GetDH() (*Dh, error) {
	data, err := cache.NewCache(1).GETDH(strconv.Itoa(int(dh.DeviceId)))
	if err != nil {
		return nil, err
	}
	x := &Dh{}
	if err := json.Unmarshal(data, x); err != nil {
		return nil, err
	}
	return x, nil
}

func (dh *Dh) SendPrivateKey(privateKey string) error {
	dh.PrivateKey = privateKey
	marshal, err := json.Marshal(dh)
	if err != nil {
		return err
	}

	if err := cache.NewCache(1).SETDH(strconv.Itoa(int(dh.DeviceId)), marshal); err != nil {
		return err
	}
	return nil
}
