package account

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/identity"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"strconv"
)

func (a *account) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	db := cockroach.GetDB()

	v := NewAuthorization(in.Username, in.Password)
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	client, err := device.GetDeviceClient()
	if err != nil {
		return nil, err
	}

	d, err := client.CreateDevice(ctx, &v1alpha1.CreateDeviceRequest{
		AccountId: strconv.Itoa(int(v.ID)),
		Ua:        in.Ua,
		Hash:      uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}

	// Creating an authorization token.
	k, err := identity.GenToken(strconv.Itoa(int(v.ID)), v.Mail, v.Username, d.DeviceId)
	if err != nil {
		return &pb.VerifyResponse{Code: "401", Reply: err.Error()}, err
	}

	return &pb.VerifyResponse{
		Code:      "200",
		Reply:     "ok",
		Id:        strconv.Itoa(int(v.ID)),
		Token:     k,
		Mail:      v.Mail,
		DeviceId:  d.DeviceId,
		PublicKey: d.PublicKey,
	}, nil
}

func (a *account) GetPublicKeyByAccountUsername(ctx context.Context, in *pb.GetPublicKeyByAccountUsernameRequest) (*pb.GetPublicKeyByAccountUsernameResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).First(&a.Actors).Error; err != nil {
		return nil, err
	}

	return &pb.GetPublicKeyByAccountUsernameResponse{
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
