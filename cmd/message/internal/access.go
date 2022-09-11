package internal

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/matrix"
	"github.com/hvxahv/hvx/microsvc"
	"gorm.io/gorm"
)

const AccountsTable = "matrices"

type Matrices struct {
	gorm.Model

	ActorId    uint   `gorm:"primaryKey;type:bigint;actor_id"`
	DeviceId   string `gorm:"type:text;device_id"`
	HomeServer string `gorm:"type:text;home_server"`
	UserId     string `gorm:"type:text;user_id"`
}

type Matrix interface {
	// Create The matrix information will be created, using the ActorId as the association.
	Create() error

	// Register for a Matrix account.
	// Verify whether the user is correct, register to matrix after successful verification,
	// and return the registration information to the client.
	Register(username, password string) (*matrix.RegisterRes, error)
}

// NewMatrices constructor for NewMatrices to create matrix data.
func NewMatrices(actorId uint, deviceId, homeServer, userId string) *Matrices {
	return &Matrices{ActorId: actorId, DeviceId: deviceId, HomeServer: homeServer, UserId: userId}
}

func (a *Matrices) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Matrices{}); err != nil {
		return errors.NewDatabaseCreate(serviceName)
	}

	if err := db.Debug().
		Table(AccountsTable).
		Create(&a).Error; err != nil {
		return err
	}
	return nil
}

// NewRegister registers to the constructor of matrix.
func NewRegister() *Matrices {
	return &Matrices{}
}

func (a *Matrices) Register(username, password string) (*matrix.RegisterRes, error) {
	var (
		ctx     = context.Background()
		address = matrix.GetRegisterAddress()
		data    = matrix.NewRegisterReq(a.DeviceId, username, password)
	)

	// Go to the account server to verify the username and password.
	v, err := clientv1.New(ctx, microsvc.AccountServiceName).Verify(username, password)
	if err != nil {
		return nil, errors.New(errors.ErrAccountVerification)
	}

	// Go to the matrix server to register an account.
	d, err := matrix.NewMatrixReq(address, data).Do()
	if err != nil {
		return nil, errors.New(errors.ErrMatrixAccountRegister)
	}
	// ERR PROCESSING...
	if d.Code != 200 {
		errors.Throw("access api error when registering to matrix server ", errors.New(string(d.Body)))
		var unmarshal matrix.RegisterErrRes
		if err := json.Unmarshal(d.Body, &unmarshal); err != nil {
			return nil, err
		} else {
			return nil, errors.Newf("%s:%s", unmarshal.Errcode, unmarshal.Error)
		}
	}

	// If the registration is successful, the returned results are processed
	// and the information is sent to the database for persistence.
	var x matrix.RegisterRes
	if err := json.Unmarshal(d.Body, &x); err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(v.ActorId)
	if err != nil {
		return nil, err
	}
	if err := NewMatrices(uint(aid), a.DeviceId, matrix.GetMatrixAddress(), x.UserId).Create(); err != nil {
		return nil, err
	}

	return &x, nil
}
