package models

import "github.com/jinzhu/gorm"

type Articles struct {
	gorm.Model
	Article 	string`gorm:"size:2000"`
	Author		string`gorm:"author"`
	Image		string`gorm:"image"`
	Status      bool`gorm:"status"`
	private		bool`gorm:"private"`
	IsComment	bool`gorm:"private"`
}

func (a *Articles) NewArticle() *Articles {
	return a
}

