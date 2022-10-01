package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
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
