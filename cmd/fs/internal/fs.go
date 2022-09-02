package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
)

const (
	FsTableName = "fs"
)

type Fs struct {
	gorm.Model

	AccountId uint   `gorm:"index;type:bigint;account_id"`
	FileName  string `gorm:"type:text;file_name"`
	Address   string `gorm:"type:text;address"`
}

func NewFs(accountId uint, fn string) *Fs {
	return &Fs{AccountId: accountId, FileName: fn}
}

type FS interface {
	Create() error
	Delete() error
	Get() (*Fs, error)
}

func NewFsCreate(accountId uint, fn, address string) *Fs {
	return &Fs{AccountId: accountId, FileName: fn, Address: address}
}

func (f *Fs) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Fs{}); err != nil {
		return err
	}

	if err := db.Debug().
		Table(FsTableName).
		Create(f).Error; err != nil {
		return err
	}
	return nil
}

func (f *Fs) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(FsTableName).
		Where("account_id =? AND file_name = ?", f.AccountId, f.FileName).
		Unscoped().
		Delete(&Fs{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *Fs) Get() (*Fs, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(FsTableName).
		Where("account_id =? AND file_name = ?", f.AccountId, f.FileName).
		First(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}
