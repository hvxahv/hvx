package account

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/policy"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

func (a *account) Verify(ctx context.Context, in *pb.NewAccountVerify) (*pb.VerifyAccountReply, error) {
	db := cockroach.GetDB()

	v := NewAuthorization(in.Username, in.Password)
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	log.Println(v)
	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}
	deviceID := uuid.New().String()

	t, err := policy.GenToken(strconv.Itoa(int(v.ID)), v.Mail, v.Username, v.Password, deviceID)
	if err != nil {
		return &pb.VerifyAccountReply{Code: "401", Reply: err.Error()}, err
	}

	d := device.NewDevices(v.ID, in.Ua, deviceID)
	if err := d.Create(); err != nil {
		return &pb.VerifyAccountReply{Code: "500", Reply: err.Error()}, err
	}

	return &pb.VerifyAccountReply{
		Code:      "200",
		Reply:     "ok",
		Id:        strconv.Itoa(int(v.ID)),
		Token:     t,
		Mail:      v.Mail,
		DeviceId:  d.Hash,
		PublicKey: d.PublicKey,
	}, nil
}

func NewAuthorization(username string, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}
