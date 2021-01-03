package app

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	Status 	string`gorm:"size:2000"`
	Author		string`gorm:"author"`
}

func NewStatus(con, author string) *Status {
	a := new(Status)
	a.Status = con
	a.Author = author
	return a
}
