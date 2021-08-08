package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

const (
	ERROR_NEW_ACCOUNT   = "FAILED TO CREATE ACCOUNTS!"
	SERVER_ERROR        = "SERVER ERROR!"
	SUCCESS_NEW_ACCOUNT = "NEW ACCOUNT OK!"
	EXISTS_MAIL         = "MAIL_EXISTS!"
	EXISTS_USERNAME     = "USERNAME_EXISTS!"
	EXISTS_ACCOUNTS     = "ACCOUNTS_EXISTS!"
)

// AccountData The object tops a user’s profile data and is targeted at GORM.
// Must be a unique key: username, email and phone.
type AccountData struct {
	gorm.Model
	Uuid       string `gorm:"type:varchar(100);uuid"`
	Username   string `gorm:"primaryKey;type:varchar(100);username;unique"`
	Password   string `gorm:"type:varchar(100);password"`
	Avatar     string `gorm:"type:varchar(100);avatar"`
	Bio        string `gorm:"type:varchar(999);bio"`
	Name       string `gorm:"type:varchar(100);name"`
	Mail       string `gorm:"primaryKey;type:varchar(100);mail;unique"`
	Phone      string `gorm:"type:varchar(100);phone"`
	Private    int32  `gorm:"private"`
	PrivateKey string `gorm:"type:varchar(3000);private_key"`
	PublicKey  string `gorm:"type:varchar(3000);public_key"`
}

// Accounts The interface defines the CRUD function for accounts.
type Accounts interface {
	// New Add a user Instantiate using the NewAccounts function.
	New() (int32, string)

	// Find This method implements the function of querying accounts.
	// It needs to accept the username to be queried through the function of the
	// instantiated object NewAccount,
	// and then return the query error and the data of the accounts structure.
	Find() (*AccountData, error)

	// Update Use the NewAccountQUD function to pass the username and
	// accept the accounts object data to update the accounts data.
	Update() error

	// Delete Pass the user name through the NewAccountQUD function to delete the user.
	Delete() error

	// Login to the account and generate token, Return token and custom error message.
	Login() (string, string, error)
}

func NewAccounts(username, password, mail string) (Accounts, error) {
	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("failed to generate public and private keys: %v", err)
		return nil, err
	}
	id := uuid.New().String()
	hash := security.GenPassword(password)

	return &AccountData{
		Uuid:       id,
		Username:   username,
		Mail:       mail,
		Password:   hash,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func NewAcctByName(name string) Accounts {
	return &AccountData{Username: name}
}

func NewAccountLogin(mail string, password string) Accounts {
	return &AccountData{Mail: mail, Password: password}
}

type acc struct {
	Username   string `gorm:"primaryKey;type:varchar(100);username;unique"`
}
func (a *AccountData) New() (int32, string) {
	d := db.GetDB()

	if err := d.AutoMigrate(&AccountData{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, SERVER_ERROR
	}

	// Check if the username and mail exist from the cache.
	mail := cache.FINDAcctMail(a.Mail)
	user := cache.ExistAcct(a.Username)
	if mail == true || user == true {
		var r string
		if mail == true {
			r = EXISTS_MAIL
		}
		if user == true {
			r = EXISTS_USERNAME
		}
		if user && mail == true {
			r = fmt.Sprintf("%s_AND_%s", EXISTS_MAIL, EXISTS_USERNAME)
		}
		return 202, r
	}

	err := d.Debug().
		Table("account_data").
		Where("username = ? ", a.Username).Or("mail = ?", a.Mail).
		First(&AccountData{}).Error

	if err != nil {
		fmt.Println("CREATE ACCOUNTS")
	}

	//// Before creating, first check whether the user exists. If it does not exist, create the user.
	//// If it does, it needs to return an error to the client to explain that the user already exists.
	//if r := d.Debug().Table("account_data").Where("username = ?", a.Username).Find(&a); r.Error != nil {
	//	fmt.Println("，没找到")
	//	if r.Error == gorm.ErrRecordNotFound {
	//		if err := d.Debug().Table("account_data").Create(&a).Error; err != nil {
	//			log.Printf("an error occurred while creating the account: %v", err)
	//			return 500, ERROR_NEW_ACCOUNT
	//		}
	//
	//		// After the user is successfully created,
	//		// the data encoded by the user's json is stored in the cache,
	//		// and the cache will never expire.
	//		ad, _ := json.Marshal(&a)
	//		if err := cache.SETAcct(a.Username, ad, 0); err != nil {
	//			log.Println(err)
	//		}
	//
	//		if err := cache.SETAcctMail(a.Mail); err != nil {
	//			log.Println(err)
	//		}
	//
	//		// TODO - Hand over to the notification server.
	//
	//		// 201 The request is successful and the server has created a new resource.
	//		return 201, SUCCESS_NEW_ACCOUNT
	//	}
	//}
	// It will not be judged so detailed in the database,
	// it just returns the error that the user has created.
	return 202, EXISTS_ACCOUNTS
}

func (a *AccountData) Find() (*AccountData, error) {
	d := db.GetDB()

	result, err := cache.GetRDB().Get(context.Background(), a.Username).Result()
	if err != nil {
		// If the cache is not found, the data will be searched from the database.
		if err := d.Debug().Table("account_data").Where("username = ?", a.Username).First(&a).Error; err != nil {
			log.Println(gorm.ErrMissingWhereClause)
			return nil, err
		}
		// The data obtained from the database is stored in the cache again.
		ad, _ := json.Marshal(&a)
		if sce := cache.SETAcct(a.Username, ad, 0); sce != nil {
			return nil, err
		}

		return a, nil
	}
	// If the cache is found, the data in the cache will be returned.
	if sre := json.Unmarshal([]byte(result), a); sre != nil {
		log.Println("Accounts failed to find user cache and parse json.")
	}
	return a, nil

}

func (a *AccountData) Update() error {
	d := db.GetDB()
	acct := &a

	if err := d.Debug().Table("account_data").Where("username = ?", a.Username).Updates(&acct).First(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	} else {
		// update data to the cache server.
		ad, _ := json.Marshal(&acct)
		if err := cache.SETAcct(a.Username, ad, 0); err != nil {
			fmt.Println("存储到 reids 失败")
			log.Println(err)
		}

	}
	return nil
}

func (a *AccountData) Delete() error {
	d := db.GetDB()
	//  Unscoped() Use gorm's Unscoped method to permanently delete data.
	if err := d.Debug().Table("account_data").Where("username = ?", a.Username).Unscoped().Delete(&a).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	}
	if err := cache.DelKey(a.Username); err != nil {
		return err
	}
	return nil
}

func (a *AccountData) Login() (string, string, error) {
	d := db.GetDB()

	var qa *AccountData
	if err := d.Debug().Table("account_data").Where("mail = ?", a.Mail).First(&qa).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(qa.Password), []byte(a.Password)); err != nil {
		return "", "", errors.Errorf("Password verification failed.")
	}
	return qa.Username, qa.Uuid, nil
}
