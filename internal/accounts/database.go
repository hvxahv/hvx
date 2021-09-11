package accounts

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// IsNotFound If the error is NOT FOUND, return TRUE otherwise return FALSE.
func IsNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}

// FoundAccount Determine whether the user exists in the database.
func FoundAccount(username, mail string) bool {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		fmt.Printf("failed to automatically create database: %v", err)
	}

	err := db.Debug().Table("accounts").Where("username = ? ", username).Or("mail = ?", mail).First(&Accounts{})
	if err != nil {
		ok := IsNotFound(err.Error)
		return ok
	}
	return false
}

// AccountIsNotFound Determine whether the user exists.
func AccountIsNotFound(username string) (*Accounts, bool) {
	db := cockroach.GetDB()

	var acct *Accounts
	r := db.Debug().Table("accounts").Where("username = ? ", username).First(&acct)
	if r.Error != nil {
		ok := IsNotFound(r.Error)
		return nil, ok
	}

	return acct, false
}

// NewAccount Create account method.
func NewAccount(acct *Accounts) error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("accounts").Create(&acct).Error; err != nil {
		return errors.Errorf("An error occurred while creating the account: %v", err)
	}

	return nil
}

// AccountLogin Log in to the account and return the account name.
func AccountLogin(mail, password string) (string, error) {
	db := cockroach.GetDB()

	var a *Accounts
	if err := db.Debug().Table("accounts").Where("mail = ?", mail).First(&a).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return "", errors.Errorf("Password verification failed.")
	}
	return a.Username, nil
}

// AccountUpdate ...
func AccountUpdate(a *Accounts) (*Accounts, error) {
	db := cockroach.GetDB()

	acct := a
	err := db.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&acct).First(&a).Error
	if err != nil {
		return nil, errors.Errorf("failed to update user: %v", err)
	}
	return acct, nil
}

// DeleteAccount Unscoped() Use gorm's Unscoped method to permanently delete data.
func DeleteAccount(name string) error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("accounts").Where("username = ?", name).Unscoped().Delete(&Accounts{}).Error; err != nil {
		return errors.Errorf("failed to delete accounts: %v", err)
	}
	return nil
}

// IsUNFo ...
func IsUNFo(name, target string) bool {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Follows{}); err != nil {
		fmt.Printf("failed to automatically create database: %v\n", err)
	}
	err := db.Debug().Table("follows").Where("name = ? AND target_name = ?", name, target).First(&Accounts{})
	if err != nil {
		ok := IsNotFound(err.Error)
		return ok
	}
	return false
}

// NewFollow ...
func NewFollow(fo *Follows) error {
	db := cockroach.GetDB()

	ok := IsUNFo(fo.Name, fo.TargetName)
	if !ok {
		return errors.Errorf("ACTOR ALREADY FOLLOWED.")
	}

	if err := db.Debug().Table("follows").Create(&fo).Error; err != nil {
		return errors.Errorf("failed to follow actor: %v", err)
	}

	if err := db.Debug().Table("accounts").Where("username = ?", fo.Name).
		Update("following", gorm.Expr("following + ?", 1)).Error; err != nil {
			return errors.Errorf("failed to increase the number of followers: %v", err)
	}

	return nil
}

func FetchAccountFollowing(fo Follows) {
	db := cockroach.GetDB()
	res := Follows{}

	if r := db.Debug().Table("follows").Where("name = ?", fo.Name).Find(&res); r.Error != nil {
		log.Printf(r.Error.Error())
	}

	fmt.Println(res.TargetName)
}

func FetchAccountFollowers(fo Follows) {
	db := cockroach.GetDB()
	res := Follows{}

	if r := db.Debug().Table("follows").Where("target_name = ?", fo.TargetName).Find(&res); r.Error != nil {
		log.Printf(r.Error.Error())
	}

	fmt.Println(res.Name)
}