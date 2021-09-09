package accounts

import (
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


// AccountIsNotFound Determine whether the user exists in the database.
func AccountIsNotFound(username, mail string) bool {
	db := cockroach.GetDB()
	err := db.Debug().Table("accounts").Where("username = ? ", username).Or("mail = ?", mail).First(&Accounts{})
	if err != nil {
		ok := IsNotFound(err.Error)
		return ok
	}
	return false
}

// AccountUserIsNotFound Determine whether the user exists.
func AccountUserIsNotFound(username string) (*Accounts, bool) {
	db := cockroach.GetDB()
	var acct *Accounts
	err := db.Debug().Table("accounts").Where("username = ? ", username).First(&acct)
	if err != nil {
		ok := IsNotFound(err.Error)
		return nil, ok
	}
	return acct, false
}

// FindAccountID Find the user ID of an account by username.
func FindAccountID(username string) (uint, error) {
	db := cockroach.GetDB()

	var acct *Accounts
	if err := db.Debug().Table("accounts").Where("username = ?",username).First(&acct).Error; err != nil {
		return 0, err
	}
	return acct.ID, nil
}

// NewAccount Create account method.
func NewAccount(acct *Accounts) error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&acct); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("accounts").Create(&acct).Error; err != nil {
		return errors.Errorf("An error occurred while creating the account: %v", err)
	}

	var aa *AccountAliases
	if err := db.AutoMigrate(&aa); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	c := &AccountAliases{
		Model:    gorm.Model{
			ID:        acct.ID,
		},
		Username: acct.Username,
	}

	if err := db.Debug().Table("account_aliases").Create(&c).Error; err != nil {
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
	if err := db.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&acct).First(&a).Error; err != nil {
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