package app

func GetAccountData(u string) *Accounts {
	a := new(Accounts)
	if db2.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		return nil
	} else {
		db2.Debug().Table("accounts").Where("username = ?", u).First(&a)
		return a
	}
}

