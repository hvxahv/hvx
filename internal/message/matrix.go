package message

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Matrices struct {
	gorm.Model

	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id;"`
	Token      string `gorm:"type:text;token"`
	HomeServer string `gorm:"type:text;home_server"`
	UserId     string `gorm:"type:text;user_id"`
	DeviceID   string `gorm:"type:text;device_id"`
}

func (m *Matrices) UpdateToken() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("matrices").Where("account_id = ?", m.AccountID).Updates(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Matrices) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Matrices{}); err != nil {
		return err
	}

	if err := db.Debug().Table("matrices").Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func NewMatrixAccesses(accountID uint, token, homeServer, userId, deviceID string) *Matrices {
	return &Matrices{AccountID: accountID, Token: token, HomeServer: homeServer, UserId: userId, DeviceID: deviceID}
}

func NewAccessUpdateToken(accountID uint, token string) *Matrices {
	return &Matrices{
		AccountID: accountID,
		Token:     token,
	}
}

type Accesses interface {
	Create() error
	UpdateToken() error
}
