package app

import "github.com/jinzhu/gorm"

type Accounts struct {
	gorm.Model
	Username	string`gorm:"username"`
	Password	string`gorm:"password"`
	Avatar		string`gorm:"avatar"`
	Bio			string`gorm:"bio"`
	Name 		string`gorm:"name"`
	EMail		string`gorm:"email"`
	Phone		int`gorm:"phone"`
	Telegram	string`gorm:"telegram"`
	Social      string`gorm:"social"`
	private     int`gorm:"private"`
}

func NewAccounts(un, p string) *Accounts {
	a := new(Accounts)
	a.Username = un
	a.Password = p
	return a
}

