package services

import (
	"hvxahv/pkg/database"
	"hvxahv/pkg/structs"
)
// GetAccountData 获取账户数据，将数据返回给调用者
func GetAccountData(u string) *structs.Accounts {
	db := database.GetMaria()
	a := new(structs.Accounts)
	if db.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		return nil
	} else {
		db.Debug().Table("accounts").Where("username = ?", u).First(&a)
		return a
	}
}

// GetActorData 获取 Actor ,
func GetActorData(u string) *structs.Accounts {
	db := database.GetMaria()
	a := new(structs.Accounts)
	db.Debug().Table("accounts").Where("username = ?", u).First(&a)

	return a
}
