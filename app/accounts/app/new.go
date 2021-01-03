package app

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	pb "hvxahv/api/kernel/v1"
)

func NewAccount(in *pb.AccountData) int {
	u := in.Username
	p := in.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密密码失败", err)
	}

	a := *NewAccounts(u, string(hash))
	db2.AutoMigrate(Accounts{})
	// If No User is Create Accounts
	if db2.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		db2.Debug().Table("accounts").Create(&a)
		db2.LogMode(true)
		if db2.NewRecord(a){
			return 500
		} else {
			return 200
		}
	} else  {
		return 202
	}
}