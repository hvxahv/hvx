package app

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	pb "hvxahv/api/kernel/v1"
)

// NewAccount 创建账户的方法
func NewAccount(in *pb.AccountData) int {
	u := in.Username
	p := in.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密密码失败", err)
	}

	a := *NewAccounts(u, string(hash))
	db2.AutoMigrate(Accounts{})
	// 账户名为唯一，如果没有这个账户名就创建账户，如果有就返回该账户已经存在
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