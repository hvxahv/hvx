package internal

import (
	"time"

	pb "github.com/hvxahv/hvx/APIs/grpc-go/account/v1alpha1"
	"github.com/hvxahv/hvx/client/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

type Public struct {
	ctx context.Context
}

type public interface {
	AccountIsExist(name string) (bool, error)
	GetActorByUsername(username string) (*pb.ActorDataResponse, error)
	CreateAccounts(username, mail, password, publicKey string) (*pb.CreateAccountResponse, error)
	Auth(username, password string) (*pb.VerifyResponse, error)
}

func NewPublic(ctx context.Context) *Public {
	return &Public{ctx: ctx}
}

func (p *Public) AccountIsExist(name string) (bool, error) {
	c, err := clientv1.New(p.ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return false, err
	}
	defer c.Close()
	exist, err := c.IsExist(p.ctx, &pb.IsExistRequest{
		Username: name,
	})
	if err != nil {
		return false, err
	}
	return exist.IsExist, nil
}

func (p *Public) GetActorByUsername(username string) (*pb.ActorDataResponse, error) {
	c, err := clientv1.New(p.ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	a, err := c.GetActorByUsername(p.ctx, &pb.GetActorByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return a, err
}

func (p *Public) CreateAccounts(username, mail, password, publicKey string) (*pb.CreateAccountResponse, error) {
	c, err := clientv1.New(p.ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.CreateAccount(p.ctx, &pb.CreateAccountRequest{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (p *Public) Auth(username, password string) (*pb.VerifyResponse, error) {
	c, err := clientv1.New(p.ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.Verify(p.ctx, &pb.VerifyRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return res, err
}
