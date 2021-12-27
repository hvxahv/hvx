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

	AccountID uint   `gorm:"primaryKey;type:bigint;accounts_id"`
	Hash      string `gorm:"type:text;hash"`
	FileType  string `gorm:"type:text;file_type"`
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

func NewSaves(accountID uint, hash string, fileType string) *Saves {
	return &Saves{AccountID: accountID, Hash: hash, FileType: fileType}
}

type Saved interface {
	Create() error
}
