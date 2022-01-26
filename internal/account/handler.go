package account

import (
	"fmt"
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/policy"
	"golang.org/x/net/context"
	"strconv"
)

func (s *server) Create(ctx context.Context, in *pb.NewCreate) (*pb.CreateReply, error) {
	if err := NewAccounts(in.Username, in.Mail, in.Password).Create(); err != nil {
		return &pb.CreateReply{Code: "202", Reply: err.Error()}, nil
	}

	_, err := NewActors(in.Username, in.PublicKey, "Person").Create()
	if err != nil {
		return &pb.CreateReply{Code: "202", Reply: err.Error()}, nil
	}

	return &pb.CreateReply{Code: "200", Reply: "ok"}, nil
}

func (s *server) Verify(ctx context.Context, in *pb.NewVerify) (*pb.VerifyReply, error) {
	verify, err := NewAuth(in.Username, in.Password).VerifyAccount()
	if err != nil {
		fmt.Println(err)
	}

	deviceID := uuid.New().String()
	token, err := policy.GenToken(verify.AccountID, verify.Mail, verify.Username, verify.Password, deviceID)
	if err != nil {
		return &pb.VerifyReply{Code: "401", Reply: err.Error()}, nil
	}

	d := device.NewDevices(verify.AccountID, in.Ua, deviceID)
	if err := d.Create(); err != nil {
		return &pb.VerifyReply{Code: "500", Reply: err.Error()}, nil
	}

	return &pb.VerifyReply{
		Code:      "200",
		Reply:     "ok",
		Token:     token,
		Mail:      verify.Mail,
		DeviceId:  deviceID,
		PublicKey: d.PublicKey,
	}, nil
}

func (s *server) GetActorByAccountUsername(ctx context.Context, in *pb.NewAccountUsername) (*pb.ActorData, error) {
	actor, err := NewActorsPreferredUsername(in.Username).GetActorByAccountUsername()
	if err != nil {
		return nil, err
	}
	return &pb.ActorData{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

func (s *server) GetAccountByUsername(ctx context.Context, in *pb.NewAccountUsername) (*pb.AccountData, error) {
	account, err := NewAccountsUsername(in.Username).GetAccountByUsername()
	if err != nil {
		return nil, err
	}
	return &pb.AccountData{
		AccountId: strconv.Itoa(int(account.ID)),
		Username:  account.Username,
		Mail:      account.Mail,
		Password:  account.Password,
		ActorId:   strconv.Itoa(int(account.ActorID)),
		IsPrivate: strconv.FormatBool(account.IsPrivate),
	}, nil
}

func (s *server) DeleteAccount(ctx context.Context, in *pb.NewAccountDelete) (*pb.DeleteAccountReply, error) {
	account, err := NewAuth(in.Username, in.Password).VerifyAccount()
	if err != nil {
		return nil, err
	}

	if err := NewAccountsID(account.AccountID).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteAccountReply{Code: "200", Reply: "ok"}, nil
}
