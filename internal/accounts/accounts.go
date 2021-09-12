package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/disism/hvxahv/internal"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

// Accounts The object tops a user’s profile data and is targeted at GORM.
// Must be a unique key: username, email and phone.
type Accounts struct {
	gorm.Model
	Username   string `gorm:"primaryKey;type:varchar(100);username;unique" validate:"required,min=4,max=16"`
	Password   string `gorm:"type:varchar(100);password" validate:"required,min=8,max=100"`
	Avatar     string `gorm:"type:varchar(999);avatar"`
	Bio        string `gorm:"type:varchar(999);bio" validate:"max=650"`
	Name       string `gorm:"type:varchar(100);name" validate:"max=16"`
	Mail       string `gorm:"index;type:varchar(100);mail;unique" validate:"required,email"`
	Phone      string `gorm:"type:varchar(100);phone"`
	IsPrivate  bool   `gorm:"type:boolean;is_private"`
	Follower   int    `gorm:"type:bigint;follower"`
	Following  int    `gorm:"type:bigint;following"`
	Friend     int    `gorm:"type:bigint;friend"`
	PrivateKey string `gorm:"type:varchar(3000);private_key"`
	PublicKey  string `gorm:"type:varchar(3000);public_key"`
}


func (a *Accounts) QueryByID() (*Accounts, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	r := db.Debug().Table("accounts").Where("id = ? ", a.ID).First(&acct)
	if r.Error != nil {
		ok := cockroach.IsNotFound(r.Error)
		if !ok {
			return acct, nil
		}
	}

	return acct, nil
}

func (a *Accounts) New() (int32, string) {
	// Check if the username and mail exist from the cache.
	mail := cache.SISAcctMail(a.Mail)
	user := cache.SISAcct(a.Username)
	if mail == true || user == true {
		var r string
		if mail == true {
			r = internal.ExistsMail
		}
		if user == true {
			r = internal.ExistsUsername
		}
		if user && mail == true {
			r = fmt.Sprintf("%s_AND_%s", internal.ExistsMail, internal.ExistsUsername)
		}
		return 202, r
	}

	// Before creating, first, check whether the user exists. If it does not exist, create the user.
	// If the account is found, it returns the user has created.
	// It will not be judged so detailed in the database. It only returns the information created by the user.
	// If more processing is required, detailed operations need to be done in the cache.
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		fmt.Printf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("accounts").Where("username = ? ", a.Username).Or("mail = ?", a.Mail).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return 400, internal.ExistsAccounts
		}
	}

	// New account.
	if err := db.Debug().Table("accounts").Create(&a).Error; err != nil {
		return 200, fmt.Sprintf("An error occurred while creating the account: %v", err)
	}

	// When the user is successfully created, the data needs to be synchronized to the cache,
	// Use JSON encoding, and the cache will never expire.
	m, _ := json.Marshal(&a)
	if err := cache.SETAcct(a.Username, m, 0); err != nil {
		return 500, "SYNC TO CACHE FAILED"
	}
	if err := cache.SETAcctMail(a.Mail); err != nil {
		return 500, "SYNC TO CACHE FAILED"
	}

	return 201, internal.SuccessNewAccount
}

func (a *Accounts) QueryByName() (*Accounts, error) {
	//r, err := cache.GetAccount(a.Username)
	//if err != nil {
	// If the cache is not found, the data will be searched from the database.

	db := cockroach.GetDB()

	var acct *Accounts
	r := db.Debug().Table("accounts").Where("username = ? ", a.Username).First(&acct)
	if r.Error != nil {
		ok := cockroach.IsNotFound(r.Error)
		if !ok {

			// The data obtained from the database is stored in the cache again.
			//ad, _ := json.Marshal(&acct)
			//if sce := cache.SETAcct(a.Username, ad, 0); sce != nil {
			//	return nil, err
			//}
			return acct, nil
		}
	}

	//return nil, errors.Errorf("THE CURRENT USER DOES NOT EXIST.")

	// If the cache is found, the data in the cache will be returned.
	//if err := json.Unmarshal([]byte(r), a); err != nil {
	//	log.Println("accounts failed to find user cache and parse json.")
	//}

	return acct, nil

}

func (a *Accounts) Update() error {
	// Password Re:encryption.
	if a.Password != "" {
		a.Password = security.GenPassword(a.Password)
	}

	db := cockroach.GetDB()

	var acct Accounts
	err := db.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&a).First(&acct).Error
	if err != nil {
		return errors.Errorf("failed to update user: %v", err)
	}

	// update data to the cache server.
	data, _ := json.Marshal(&acct)
	if err := cache.UPDATEAcct(a.Username, data); err != nil {
		return err
	}

	if a.Mail != "" {
		if err := cache.UPDATEMail(a.Mail, acct.Mail); err != nil {
			return err
		}
	}

	return nil
}

func (a *Accounts) Delete() error {
	auth := NewAccountAuth(a.Mail, a.Password)

	name, err := auth.Login()
	if err != nil {
		return errors.Errorf("Verification failed, account cannot be deleted: %v", err)
	}

	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("username = ?", name).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return errors.Errorf("failed to delete accounts: %v", err)
	}

	if err := cache.DELKey(name); err != nil {
		return err
	}
	if err := cache.DELAcctMail(a.Mail); err != nil {
		return err
	}

	return nil
}

// Account The interface defines the CRUD function for accounts.
type Account interface {
	// New Add a user Instantiate using the NewAccounts function.
	New() (int32, string)

	// QueryByName This method implements the function of querying accounts.
	// It needs to accept the username to be queried through the function of the
	// instantiated object NewAccount,
	// and then return the query error and the data of the accounts structure.
	QueryByName() (*Accounts, error)

	QueryByID() (*Accounts, error)

	// Update Use the NewAccountQUD function to pass the username and
	// accept the accounts object data to update the accounts data.
	Update() error

	// Delete Pass the user name through the NewAccountQUD function to delete the user.
	Delete() error
}

func NewAccounts(username, password, mail string) (Account, error) {
	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("failed to generate public and private keys: %v", err)
		return nil, err
	}

	hash := security.GenPassword(password)

	acct := &Accounts{
		Username:   username,
		Mail:       mail,
		Password:   hash,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	v := validator.New()
	if err := v.Struct(*acct); err != nil {
		return nil, err
	}
	return acct, nil
}

func NewAccountByName(name string) Account {
	return &Accounts{Username: name}
}

func NewDelete(mail, password, username string) Account {
	return &Accounts{Username: username, Mail: mail, Password: password}
}

func NewAccountByID(id uint) Account {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}
