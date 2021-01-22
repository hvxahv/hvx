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
	Private     int`gorm:"private"`
	PrivateKey  string`gorm:"private_key"`
	PublicKey	string`gorm:"public_key"`
}

func NewAccounts(un, p string) *Accounts {
	a := new(Accounts)
	a.Username = un
	a.Password = p
	return a
}

