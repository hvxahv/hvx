package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/auth"
)

type Auth interface {
	SetAuthPublicKey(accountId int64, publicKey string) (*pb.SetPublicKeyResponse, error)
}

func (svc *Svc) SetAuthPublicKey(accountId int64, publicKey string) (*pb.SetPublicKeyResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	k, err := pb.NewAuthClient(c.Conn).SetPublicKey(svc.ctx, &pb.SetPublicKeyRequest{
		AccountId: accountId,
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}
	return k, err
}
