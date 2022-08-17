package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/APIs/v1alpha1/device"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

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

	// AddDevice Add a device to the device system.
	AddDevice(accountId, ua string) (*device.CreateResponse, error)
}

func (a *authorization) Authorization(username, password string) (*account.VerifyResponse, error) {
	c, err := clientv1.New(a.ctx,
		microsvc.NewGRPCAddress("account").Get(),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	v, err := account.NewAccountsClient(c.Conn).Verify(a.ctx, &account.VerifyRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return v, err
}

func (a *authorization) AddDevice(accountId, ua string) (*device.CreateResponse, error) {
	c, err := clientv1.New(a.ctx,
		microsvc.NewGRPCAddress("device").Get(),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	d, err := device.NewDevicesClient(c.Conn).Create(a.ctx, &device.CreateRequest{
		AccountId: accountId,
		Ua:        ua,
	})
	if err != nil {
		return nil, err
	}
	return d, nil
}
