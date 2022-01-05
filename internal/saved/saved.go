package saved

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"log"
)

// Change the upload logic.
// The upload should be encrypted in the client and uploaded to the IPFS API,
// and then the returned hash should be submitted to the hvxahv server for storage.
// hvxahv should not store the user's file key,
// only the saved hash value is used to format the storage list.

type Saves struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey" json:"ID,string"`
	AccountID uint   `gorm:"primaryKey;type:bigint;accounts_id" json:"account_id,string"`
	Name      string `gorm:"type:text;name"`
	Hash      string `gorm:"type:text;hash"`
	FileType  string `gorm:"type:text;file_type"`
}

func (s *Saves) GetSaves() (*[]Saves, error) {
	db := cockroach.GetDB()
	var saves []Saves
	if err := db.Debug().Table("saves").Where("account_id = ?", s.AccountID).Find(&saves).Error; err != nil {
		return nil, err
	}
	return &saves, nil
}

func (s *Saves) GetSavedByID() (*Saves, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("saves").Where("id = ?", s.ID).First(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Saves) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Saves{}); err != nil {
		log.Println(err)
		return err
	}
	if err := db.Debug().Table("saves").Create(&s).Error; err != nil {
		return err
	}
	return nil
}

func NewSaves(accountID uint, name, hash string, fileType string) *Saves {
	return &Saves{AccountID: accountID, Name: name, Hash: hash, FileType: fileType}
}

func NewSavesID(id uint) *Saves {
	return &Saves{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewSavesByAccountID(accountID uint) *Saves {
	return &Saves{AccountID: accountID}
}

type Saved interface {
	GetSaves() (*[]Saves, error)
	Create() error
	GetSavedByID() (*Saves, error)
}
