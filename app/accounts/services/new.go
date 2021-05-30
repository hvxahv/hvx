package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/internal/accounts"
	"hvxahv/pkg/maria"
	"hvxahv/pkg/utils"
	"log"
)

// NewAccount 创建账户的方法
func NewAccount(in *pb.AccountData) int {
	u := in.Username
	p := in.Password
	// 进行密码加密返回加密后的 hash， 需要转化成字符串
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密密码失败", err)
	}
	private, public, err := GenRasKey()
	if err != nil {
		log.Println(err)
	}

	a := &accounts.Accounts{
		Username:   u,
		Password:   string(hash),
		PrivateKey: private,
		PublicKey:  public,
	}

	db := maria.GetMaria()
	// 自动创建数据库 table
	db.AutoMigrate(&accounts.Accounts{})
	// 账户名为唯一，如果没有这个账户名就创建账户，如果有就返回该账户已经存在
	if db.Debug().Table("accounts").Where("username = ?", u).First(&a).RecordNotFound() {
		db.Debug().Table("accounts").Create(&a)
		db.LogMode(true)
		if db.NewRecord(a) {
			return 500
		} else {
			return 200
		}
	} else {
		return 202
	}
}

// GenRasKey 该方法调用 utils.GenerateKey 包生成 2068 位的 ras key 返回公钥和私钥
func GenRasKey() (string, string, error) {
	privateKey, publicKey, err := utils.GenerateKey(2048)
	if err != nil {
		fmt.Printf("Generate key is error: %s", err)
	}

	private := utils.EncodePrivateKey(privateKey)

	public, err := utils.EncodePublicKey(publicKey)
	if err != nil {
		fmt.Println("Encode Public Key is error: ", err)
	}

	return string(private), string(public), err
}
