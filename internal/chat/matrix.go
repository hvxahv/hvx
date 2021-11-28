package chat

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type MatrixAccesses struct {
	gorm.Model
	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id;"`
	Token      string `gorm:"type:text;token"`
	HomeServer string `gorm:"type:text;home_server"`
	UserId     string `gorm:"type:text;user_id"`
	DeviceID   string `gorm:"type:text;device_id"`
}

func (m *MatrixAccesses) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&MatrixAccesses{}); err != nil {
		return err
	}

	if err := db.Debug().Table("matrix_accesses").Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func NewMatrixAccess(accountID uint, token, homeServer, userId, deviceID string) *MatrixAccesses {
	return &MatrixAccesses{AccountID: accountID, Token: token, HomeServer: homeServer, UserId: userId, DeviceID: deviceID}
}

type MatrixAccess interface {
	Create() error
}
