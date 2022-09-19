/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 */

package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"gorm.io/gorm"
)

const (
	SavesTable = "saves"
)

// Change the upload logic. The upload should be encrypted
// in the cfg and uploaded to the  IPFS API, and then
// the returned cid should be submitted to the hvxahv
// server for storage. hvxahv should not store the user's
// file key, only the saved cid value is used to format
// the storage list.

// Saves a file to the database.
type Saves struct {
	gorm.Model

	AccountId uint `gorm:"primaryKey;type:bigint;account_id"`

	Name      string `gorm:"type:text;name"`
	Comment   string `gorm:"type:text;comment"`
	Cid       string `gorm:"type:text;cid"`
	Types     string `gorm:"type:text;types"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

type Saved interface {
	Create() error
	GetSaved() (*Saves, error)
	GetSaves() ([]*Saves, error)
	//EditSaved(id, accountId uint) error
	//DeleteSave() error
	//DeleteSaves() error
}

type Editor interface {
	EditSavedName(name string) *Saves
	EditSavedComment(comment string) *Saves
}

func NewSaves(accountID uint, name, comment, cid, types string, isPrivate bool) *Saves {
	return &Saves{
		AccountId: accountID,
		Name:      name,
		Comment:   comment,
		Cid:       cid,
		Types:     types,
		IsPrivate: isPrivate,
	}
}

func (s *Saves) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Saves{}); err != nil {
		return errors.NewDatabaseCreate(serviceName)
	}

	if err := db.Debug().
		Table(SavesTable).
		Create(&s).
		Error; err != nil {
		return err
	}
	return nil
}

func NewSavesId(savesId uint) *Saves {
	return &Saves{
		Model: gorm.Model{
			ID: savesId,
		},
	}
}

func (s *Saves) GetSaved() (*Saves, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(SavesTable).
		Where("id = ?", s.ID).
		First(&s).
		Error; err != nil {
		return nil, err
	}
	return s, nil
}

func NewSavesAccountId(accountID uint) *Saves {
	return &Saves{
		AccountId: accountID,
	}
}

func (s *Saves) GetSaves() ([]*Saves, error) {
	db := cockroach.GetDB()
	var saves []*Saves
	if err := db.Debug().
		Table(SavesTable).
		Where("account_id = ?", s.AccountId).
		Find(&saves).
		Error; err != nil {
		return nil, err
	}
	return saves, nil
}

func (s *Saves) EditSaved(id, accountId uint) error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SavesTable).
		Where("id = ? AND account_id = ?", id, accountId).
		Updates(&s).
		Error; err != nil {
		return err
	}
	return nil
}

func NewSavesDelete(savedId, accountId uint) *Saves {
	return &Saves{
		Model: gorm.Model{
			ID: savedId,
		},
		AccountId: accountId,
	}
}

func (s *Saves) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SavesTable).
		Where("id = ? AND account_id = ?", s.ID, s.AccountId).
		Unscoped().
		Delete(&Saves{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (s *Saves) DeleteSaves() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SavesTable).
		Where("account_id = ?", s.AccountId).
		Unscoped().
		Delete(&Saves{}).
		Error; err != nil {
		return err
	}
	return nil
}
