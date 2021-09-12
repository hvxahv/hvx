package articles

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
)

func NewArticle(a *Articles) error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Articles{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("articles").Create(&a); err != nil {
		return errors.Errorf("failed to create article: %v", err)
	}
	return nil
}
