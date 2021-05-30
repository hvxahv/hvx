package inbox

import "github.com/jinzhu/gorm"

type Articles struct {
	gorm.Model
	Content     string `gorm:"size:2000"`
	Author      string `gorm:"author"`
	Image       string `gorm:"image"`
	ContentType string `gorm:"content_type"`
	private     bool   `gorm:"private"`
	IsComment   bool   `gorm:"private"`
}

func (a *Articles) NewArticle() *Articles {
	return a
}
