package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/account"
)

type Accounts interface {
	IsExistAccount(username string) (*pb.IsExistResponse, error)
	GetAccountByUsername(username string) (*pb.GetByUsernameResponse, error)
	CreateAccount(username, mail, password, publicKey string) (*pb.CreateResponse, error)
	GetPrivateKey(accountId int64) (*pb.GetPrivateKeyResponse, error)
	Verify(username, password string) (*pb.VerifyResponse, error)
}

func (svc *Svc) IsExistAccount(username string) (*pb.IsExistResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	accounts, err := pb.NewAccountsClient(c.Conn).IsExist(svc.ctx, &pb.IsExistRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (svc *Svc) GetAccountByUsername(username string) (*pb.GetByUsernameResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	accounts, err := pb.NewAccountsClient(c.Conn).GetByUsername(svc.ctx, &pb.GetByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (svc *Svc) CreateAccount(username, mail, password, publicKey string) (*pb.CreateResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	accounts, err := pb.NewAccountsClient(c.Conn).Create(svc.ctx, &pb.CreateRequest{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (svc *Svc) GetPrivateKey(accountId int64) (*pb.GetPrivateKeyResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	accounts, err := pb.NewAccountsClient(c.Conn).GetPrivateKey(svc.ctx, &pb.GetPrivateKeyRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (svc *Svc) Verify(username, password string) (*pb.VerifyResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	v, err := pb.NewAccountsClient(c.Conn).Verify(svc.ctx, &pb.VerifyRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}
