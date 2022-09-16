package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/APIs/v1alpha1/device"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Auths struct {
	gorm.Model

	AccountId uint   `gorm:"primaryKey;bigint;account_id"`
	PublicKey string `gorm:"type:text;public_key"`
}

func NewAuths(accountId uint, publicKey string) *Auths {
	return &Auths{AccountId: accountId, PublicKey: publicKey}
}

type authorization struct {
	ctx context.Context
}

func NewAuthorization(ctx context.Context) *authorization {
	return &authorization{ctx: ctx}
}

type AuthorizationHandler interface {
	// Authorization Connect to the account system to verify the username and first name password and
	// return the authorization result.
	Authorization(username, password string) (*account.VerifyResponse, error)

	SetPublicKey(accountId uint, publicKey string) error
	// AddDevice Add a device to the device system.
	AddDevice(accountId int64, ua string) (*device.CreateResponse, error)
}

func (a *authorization) Authorization(username, password string) (*account.VerifyResponse, error) {
	verify, err := clientv1.New(context.Background(), microsvc.AccountServiceName).Verify(username, password)
	if err != nil {
		errors.Throw("error while connecting to the account server for authentication in public service.", err)
		return nil, err
	}
	if err != nil {
		return nil, errors.New(errors.ErrAccountVerification)
	}
	return verify, nil
}

func (a *authorization) SetPublicKey(accountId uint, publicKey string) error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Auths{}); err != nil {
		return err
	}
	auths := NewAuths(accountId, publicKey)
	if err := db.Debug().
		Table("auths").
		Create(&auths).
		Error; err != nil {
		return err
	}
	return nil
}

func (a *authorization) AddDevice(accountId int64, ua string) (*device.CreateResponse, error) {
	add, err := clientv1.New(context.Background(), microsvc.DeviceServiceName).AddDevice(accountId, ua)
	if err != nil {
		return nil, err
	}
	if err != nil {
		errors.Throw("error occurred while connecting to the device server in public service.", err)
		return nil, errors.NewFailedToConnect(microsvc.DeviceServiceName)
	}

	return add, nil
}
