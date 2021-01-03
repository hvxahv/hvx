package app

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Article 	string`gorm:"size:2000"`
	Author		string`gorm:"author"`
}

func NewArticle(con, author string) *Article {
	a := new(Article)
	a.Article = con
	a.Author = author
	return a
}
