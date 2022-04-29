package account

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/hvxahv/hvx/pkg/identity"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"strconv"
)

func (a *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	db := cockroach.GetDB()

	v := NewAuthorization(in.Username, in.Password)
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	// TODO - Added to device management.

	deviceID := ""
	publicKey := ""
	// Creating an authorization token.
	k, err := identity.GenToken(strconv.Itoa(int(v.ID)), v.Mail, v.Username, deviceID)
	if err != nil {
		return &pb.VerifyResponse{Code: "401", Reply: err.Error()}, err
	}

	return &pb.VerifyResponse{
		Code:      "200",
		Reply:     "ok",
		Id:        strconv.Itoa(int(v.ID)),
		Token:     k,
		Mail:      v.Mail,
		DeviceId:  deviceID,
		PublicKey: publicKey,
	}, nil
}

func (a *server) GetPublicKeyByAccountUsername(ctx context.Context, in *pb.GetPublicKeyByAccountUsernameRequest) (*pb.GetPublicKeyByAccountUsernameResponse, error) {
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
