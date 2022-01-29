package account

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

type Accounts struct {
	gorm.Model

	Username  string `gorm:"primaryKey;type:text;preferredUsername;" validate:"required,min=4,max=16"`
	Mail      string `gorm:"index;type:text;mail;unique" validate:"required,email"`
	Password  string `gorm:"type:text;password" validate:"required,min=8,max=100"`
	ActorID   uint   `gorm:"type:bigint;actor_id"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

func (a *account) IsExist(ctx context.Context, in *pb.NewAccountUsername) (*pb.IsAccountExistReply, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ? ", in.Username).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return &pb.IsAccountExistReply{IsExist: ok}, nil
	}
	return &pb.IsAccountExistReply{IsExist: false}, nil
}

func (a *account) Create(ctx context.Context, in *pb.NewAccountCreate) (*pb.AccountReply, error) {
	if err := validator.New().Struct(in); err != nil {
		return nil, errors.New("FAILED_TO_VALIDATOR")
	}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Actors{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACTOR_DATABASE")
	}

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_ACCOUNT_DATABASE")
	}

	// Check if the username and mail is exist.
	if err := db.Debug().Table("accounts").Where("username = ? ", in.Username).Or("mail = ?", in.Mail).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return nil, errors.New("THE_USERNAME_OR_MAIL_ALREADY_EXISTS")
		}
	}

	// Create the actor.
	n := NewActors(in.Username, in.PublicKey, "Person")
	if err := db.Debug().Table("actors").Create(&n).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACTOR")
	}

	// Create the account.
	v := NewAccounts(n.ID, in.Username, in.Mail, in.Password)
	if err := db.Debug().Table("accounts").Create(&v).Error; err != nil {
		return nil, errors.Errorf("FAILED_TO_CREATE_ACCOUNT")
	}

	return &pb.AccountReply{Code: "200", Reply: "ok"}, nil
}

func (a *account) Delete(ctx context.Context, in *pb.NewAccountDelete) (*pb.AccountReply, error) {
	v := NewAuthorization(in.Username, in.Password)

	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&v).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(in.Password)); err != nil {
		return nil, errors.Errorf("PASSWORD_VERIFICATION_FAILED")
	}

	if err := db.Debug().Table("actors").Where("id = ?", v.ActorID).Unscoped().Delete(&Actors{}).Error; err != nil {
		return nil, err
	}

	if err := db.Debug().Table("accounts").Where("id = ?", v.ID).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return nil, err
	}

	return &pb.AccountReply{Code: "200", Reply: "ok"}, nil
}

func (a *account) EditUsername(ctx context.Context, in *pb.NewEditAccountUsername) (*pb.AccountReply, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	exist, err := a.IsExist(ctx, &pb.NewAccountUsername{Username: in.Username})
	if err != nil {
		return nil, err
	}

	// If the username is Exist, return error.
	if !exist.IsExist {
		return &pb.AccountReply{Code: "401", Reply: "THE_USERNAME_ALREADY_EXISTS"}, nil
	}

	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("id = ?", uint(id)).First(&a.Accounts).Update("username", in.Username).Error; err != nil {
		return &pb.AccountReply{Code: "500", Reply: err.Error()}, err
	}

	address := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), in.Username)
	inbox := fmt.Sprintf("%s/inbox", address)
	if err := db.Debug().Table("actors").Where("id = ?", a.Accounts.ActorID).
		Update("preferred_username", in.Username).
		Update("inbox", inbox).
		Update("address", address).Error; err != nil {
		return &pb.AccountReply{Code: "500", Reply: err.Error()}, err
	}

	return &pb.AccountReply{Code: "200", Reply: "ok"}, nil
}

func (a *account) EditPassword(ctx context.Context, in *pb.NewEditAccountPassword) (*pb.AccountReply, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("id = ?", id).Update("password", hash).Error; err != nil {
		return nil, err
	}

	// Because the password is reset, all logged-in devices should be deleted
	client, err := device.NewDeviceClient()
	if err != nil {
		return nil, err
	}
	d := &v1alpha1.NewDeviceAccountID{AccountId: in.Id}
	reply, err := client.DeleteAllByAccountID(ctx, d)
	if err != nil {
		return nil, err
	}

	return &pb.AccountReply{Code: "200", Reply: reply.Reply}, nil
}

func (a *account) EditMail(ctx context.Context, in *pb.NewEditAccountMail) (*pb.AccountReply, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("id = ?", id).Update("mail", in.Mail).Error; err != nil {
		return nil, err
	}

	return &pb.AccountReply{Code: "200", Reply: "ok"}, nil
}

func (a *account) GetAccountByUsername(ctx context.Context, in *pb.NewAccountUsername) (*pb.AccountData, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Where("username = ?", in.Username).First(&a.Accounts).Error; err != nil {
		return nil, err
	}

	return &pb.AccountData{
		AccountId: strconv.Itoa(int(a.Accounts.ID)),
		Username:  a.Accounts.Username,
		Mail:      a.Accounts.Mail,
		Password:  a.Accounts.Password,
		ActorId:   strconv.Itoa(int(a.Accounts.ActorID)),
		IsPrivate: strconv.FormatBool(a.Accounts.IsPrivate),
	}, nil
}

func NewAccountsID(id uint) *Accounts {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewAccounts(actorID uint, username, mail, password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: string(hash),
		ActorID:  actorID,
	}
}
