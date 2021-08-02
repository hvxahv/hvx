package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/disism/hvxahv/pkg/redis"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

const (
	ERROR_NEW_ACCOUNT   = "FAILED TO CREATE USER!"
	SERVER_ERROR        = "SERVER ERROR!"
	SUCCESS_NEW_ACCOUNT = "NEW ACCOUNT OK!"
	EXISTS_ACCOUNT      = "ACCOUNT ALREADY EXISTS!"
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
	EMail      string `gorm:"type:varchar(100);email;unique"`
	Phone      string `gorm:"type:varchar(100);phone;unique"`
	Private    int32  `gorm:"private"`
	PrivateKey string `gorm:"type:varchar(3000);private_key"`
	PublicKey  string `gorm:"type:varchar(3000);public_key"`
}

// Accounts The interface defines the CRUD function for accounts.
type Accounts interface {
	// New Add a user Instantiate using the NewAccounts function.
	New() (int32, string)

	// Query This method implements the function of querying accounts.
	// It needs to accept the username to be queried through the function of the
	// instantiated object NewAccountQUD,
	// and then return the query error and the data of the accounts structure.
	Query() (*AccountData, error)

	// Update Use the NewAccountQUD function to pass the username and
	// accept the accounts object data to update the accounts data.
	Update() error

	// Delete Pass the user name through the NewAccountQUD function to delete the user.
	Delete() error

	// Login to the account and generate token, Return token and custom error message.
	Login() (string, error)
}

func NewAccounts(username, password string) (Accounts, error) {
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
		Password:   hash,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func NewAcctFuncByName(name string) Accounts {
	return &AccountData{Username: name}
}

func NewAccountLogin(name string, password string) Accounts {
	return &AccountData{Username: name, Password: password}
}

func (a *AccountData) New() (int32, string) {
	d := db.GetDB()

	if err := d.AutoMigrate(&AccountData{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, SERVER_ERROR
	}

	acct := &a

	if redis.ExistKey(a.Username) {
		return 202, EXISTS_ACCOUNT
	}

	// Before creating, first check whether the user exists. If it does not exist, create the user.
	// If it does, it needs to return an error to the client to explain that the user already exists.
	if r := d.Debug().Table("account_data").Where("username = ? ", a.Username).First(&acct); r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			if err := d.Debug().Table("account_data").Create(&acct).Error; err != nil {
				log.Printf("an error occurred while creating the account: %v", err)
				return 500, ERROR_NEW_ACCOUNT
			}

			// After the user is successfully created,
			// the data encoded by the user's json is stored in the cache,
			// and the cache will never expire.
			ad, _ := json.Marshal(&a)
			if err := redis.SetJsonData(a.Username, ad, 0); err != nil {
				fmt.Println(err)
			}

			// Notify the telegram bot that a new user has been added.
			//go func() {
			//	b := bot.NewBot(1, fmt.Sprintf("Added a user: %s", a.Name))
			//	if err := b.Send(); err != nil {
			//		log.Println(err)
			//	}
			//}()

			// 201 The request is successful and the server has created a new resource.
			return 201, SUCCESS_NEW_ACCOUNT
		}
	}

	ad, _ := json.Marshal(&a)
	if err := redis.SetJsonData(a.Username, ad, 0); err != nil {
		fmt.Println(err)
	}
	return 202, EXISTS_ACCOUNT
}

func (a *AccountData) Query() (*AccountData, error) {
	d := db.GetDB()

	result, err := redis.GetRDB().Get(context.Background(), a.Username).Result()
	if err != nil {
		if err := d.Debug().Table("account_data").Where("username = ?", a.Username).First(&a).Error; err != nil {
			log.Println(gorm.ErrMissingWhereClause)
			return nil, err
		}
		// The data obtained from the database is stored in the cache again.
		go func() {
			ad, _ := json.Marshal(&a)
			_, err := redis.GetRDB().Set(context.Background(), a.Username, ad, 0).Result()
			if err != nil {
				log.Println("Failed to store to cache:", err)
			}
		}()

		return a, nil
	}
	if err := json.Unmarshal([]byte(result), a); err != nil {
		log.Println("Accounts failed to find user cache and parse json.")
		// TODO - If the cache decoding fails, you should go to the database to find the data and return.
	}
	return a, nil

}

func (a *AccountData) Update() error {
	d := db.GetDB()
	acct := &a

	if err := d.Debug().Table("account_data").Where("username = ?", a.Username).Updates(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
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
	return nil
}

func (a *AccountData) Login() (string, error) {
	d := db.GetDB()

	var qa *AccountData
	if err := d.Debug().Table("account_data").Where("username = ?", a.Username).First(&qa).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(qa.Password), []byte(a.Password)); err != nil {
		return "", errors.Errorf("Password verification failed.")
	}
	token, err := security.GenToken(a.Uuid, a.Username)
	if err != nil {
		log.Println("Token generation failed!")
	}
	return token, nil
}
