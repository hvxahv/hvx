package account

import (
	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	v1alpha "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/hvxahv/hvx/pkg/conv"
	"github.com/hvxahv/hvx/pkg/identity"
	"github.com/hvxahv/hvx/pkg/v"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func (a *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	db := cockroach.GetDB()

	auth := NewAuthorization(in.Username, in.Password)
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&auth).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	conn, err := grpc.DialContext(ctx, v.GetGRPCServiceAddress("device"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := v1alpha.NewDevicesClient(conn)
	device, err := client.CreateDevice(ctx, &v1alpha.CreateDeviceRequest{
		AccountId: conv.UintToString(auth.ID),
		Ua:        in.Ua,
		Hash:      uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}

	// Creating an authorization token.
	k, err := identity.GenToken(strconv.Itoa(int(auth.ID)), auth.Mail, auth.Username, device.DeviceId)
	if err != nil {
		return &pb.VerifyResponse{Code: "401", Reply: err.Error()}, err
	}

	return &pb.VerifyResponse{
		Code:      "200",
		Reply:     "ok",
		Id:        strconv.Itoa(int(auth.ID)),
		Token:     k,
		Mail:      auth.Mail,
		DeviceId:  device.DeviceId,
		PublicKey: device.PublicKey,
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
