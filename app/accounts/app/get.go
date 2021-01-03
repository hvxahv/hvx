package app

import "log"

func GetAccountData(u string) *Accounts {
	log.Println("---------- 得到的用户名 ",u)
	a := new(Accounts)
	if db2.Debug().Table("account").Where("username = ?", u).First(&a).RecordNotFound() {
		return nil
	} else {
		db2.Debug().Table("account").Where("username = ?", u).First(&a)
		return a

	}
	return nil
}

