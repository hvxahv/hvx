package account

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/policy"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"strconv"
)

func (a *account) Verify(ctx context.Context, in *pb.NewAccountVerify) (*pb.VerifyAccountReply, error) {
	db := cockroach.GetDB()

	v := NewAuthorization(in.Username, in.Password)
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}
	hash := uuid.New().String()

	// Creating an authorization token.
	k, err := policy.GenToken(strconv.Itoa(int(v.ID)), v.Mail, v.Username, v.Password, hash)
	if err != nil {
		return &pb.VerifyAccountReply{Code: "401", Reply: err.Error()}, err
	}

	client, err := device.NewDeviceClient()
	if err != nil {
		return nil, err
	}

	d := &v1alpha1.NewDeviceCreate{
		AccountId: strconv.Itoa(int(v.ID)),
		Ua:        in.Ua,
		Hash:      hash,
	}
	create, err := client.Create(ctx, d)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyAccountReply{
		Code:       "200",
		Reply:      "ok",
		Id:         strconv.Itoa(int(v.ID)),
		Token:      k,
		Mail:       v.Mail,
		DeviceHash: hash,
		PublicKey:  create.PublicKey,
	}, nil
}

func (a *account) GetPublicKeyByAccountUsername(ctx context.Context, in *pb.NewAccountUsername) (*pb.GetPublicKeyReply, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).First(&a.Actors).Error; err != nil {
		return nil, err
	}

	return &pb.GetPublicKeyReply{
		Code:      "200",
		PublicKey: a.Actors.PublicKey,
	}, nil
}

func NewAuthorization(username string, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}
