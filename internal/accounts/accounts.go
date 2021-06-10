package accounts

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hvxahv/pkg/db"
	"hvxahv/pkg/encrypt"
	"log"
)

// accounts The object tops a user’s profile data and is targeted at GORM.
type accounts struct {
	gorm.Model
	Uuid           string `gorm:"uuid"`
	Username       string `gorm:"username"`
	Password       string `gorm:"password"`
	Avatar         string `gorm:"avatar"`
	Bio            string `gorm:"bio"`
	Name           string `gorm:"name"`
	EMail          string `gorm:"email"`
	Phone          int    `gorm:"phone"`
	Telegram       string `gorm:"telegram"`
	Social         string `gorm:"social"`
	Private        int    `gorm:"private"`
	NSFW           int    `gorm:"nsfw"`
	PrivateKey     string `gorm:"private_key"`
	PublicKey      string `gorm:"public_key"`
	FollowingCount int    `gorm:"following_count"`
	FollowerCount  int    `gorm:"follower_count"`
}

// Accounts The interface defines the CRUD function for accounts.
type Accounts interface {
	// New Add a user Instantiate using the NewAccounts function.
	New() error
	// Query This method implements the function of querying accounts.
	// It needs to accept the username to be queried through the function of the
	// instantiated object NewAccountQUD,
	// and then return the query error and the data of the accounts structure.
	Query() (*accounts, error)
	// Update Use the NewAccountQUD function to pass the username and
	// accept the accounts object data to update the accounts data.
	Update() error
	// Delete Pass the user name through the NewAccountQUD function to delete the user.
	Delete() error
}

// NewAccounts Instantiate accounts object.
func NewAccounts(
	username string,
	password string,
	avatar string,
	bio string,
	name string,
	EMail string,
	phone int,
	telegram string,
	social string,
	private int,
	NSFW int,
	followingCount int,
	followerCount int,
) Accounts {
	privateKey, publicKey, err := encrypt.GenRSA()
	if err != nil {
		log.Printf("Failed to generate public and private keys: %v", err)
	}

	id := uuid.New().String()
	return &accounts{
		Uuid:           id,
		Username:       username,
		Password:       password,
		Avatar:         avatar,
		Bio:            bio,
		Name:           name,
		EMail:          EMail,
		Phone:          phone,
		Telegram:       telegram,
		Social:         social,
		Private:        private,
		NSFW:           NSFW,
		PrivateKey:     privateKey,
		PublicKey:      publicKey,
		FollowingCount: followingCount,
		FollowerCount:  followerCount,
	}
}

// NewAccountByName The accounts object will be instantiated by the username
// for deleting, querying and modifying methods.
func NewAccountByName(username string) *accounts {
	return &accounts{Username: username}
}

func (a *accounts) New() error {
	d := db.GetDB()
	if err := d.AutoMigrate(&accounts{}); err != nil {
		return nil
	}

	acct := &accounts{
		Username:       a.Username,
		Password:       a.Password,
		Avatar:         a.Avatar,
		Bio:            a.Bio,
		Name:           a.Name,
		EMail:          a.EMail,
		Phone:          a.Phone,
		Telegram:       a.Telegram,
		Social:         a.Social,
		Private:        a.Private,
		NSFW:           a.NSFW,
		PrivateKey:     a.PrivateKey,
		PublicKey:      a.PublicKey,
		FollowingCount: a.FollowingCount,
		FollowerCount:  a.FollowerCount,
	}

	if err := d.Debug().Table("accounts").Create(&acct).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *accounts) Query() (*accounts, error) {
	d := db.GetDB()
	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).First(&a).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return nil, err
	}

	return a, nil
}

func (a *accounts) Update() error {
	d := db.GetDB()
	acct := &accounts{
		Username:       a.Username,
		Password:       a.Password,
		Avatar:         a.Avatar,
		Bio:            a.Bio,
		Name:           a.Name,
		EMail:          a.EMail,
		Phone:          a.Phone,
		Telegram:       a.Telegram,
		Social:         a.Social,
		Private:        a.Private,
		NSFW:           a.NSFW,
		PrivateKey:     a.PrivateKey,
		PublicKey:      a.PublicKey,
		FollowingCount: a.FollowingCount,
		FollowerCount:  a.FollowerCount,
	}


	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	}
	return nil
}

func (a *accounts) Delete() error {
	d := db.GetDB()
	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).Delete(&a).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	}
	return nil
}
