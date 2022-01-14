package article

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Attachments struct {
	gorm.Model

	Name     string `gorm:"type:text;name"`
	FileType string `gorm:"type:text;file_type"`
	URL      string `gorm:"type:text;url"`
}

func (a *Attachments) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Attachments{}); err != nil {
		return err
	}
	if err := db.Debug().Where("attachments").Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func NewAttachments(name string, fileType string, URL string) *Attachments {
	return &Attachments{FileType: fileType, URL: URL, Name: name}
}

type Attachment interface {
	Create() error
}
