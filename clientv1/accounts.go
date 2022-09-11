package clientv1

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/account"
)

type Accounts interface {
	GetPrivateKey(accountId string) (*account.GetPrivateKeyResponse, error)
}

func (svc *Svc) GetPrivateKey(accountId string) (*account.GetPrivateKeyResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	accounts, err := account.NewAccountsClient(c.Conn).GetPrivateKey(svc.ctx, &account.GetPrivateKeyRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
