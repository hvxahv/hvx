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
	Uuid       string `gorm:"uuid;->:false;<-:create"`
	Username   string `gorm:"type:varchar(100);username;unique"`
	Password   string `gorm:"type:varchar(100);password"`
	Avatar     string `gorm:"type:varchar(100);avatar"`
	Bio        string `gorm:"type:varchar(999);bio"`
	Name       string `gorm:"type:varchar(100);name"`
	EMail      string `gorm:"type:varchar(100);email"`
	Phone      string `gorm:"type:varchar(100);phone"`
	Telegram   string `gorm:"type:varchar(100);telegram"`
	Private    int    `gorm:"private"`
	PrivateKey string `gorm:"type:varchar(3000);private_key;->:false;<-:create;<-:update"`
	PublicKey  string `gorm:"type:varchar(3000);public_key"`
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
	phone string,
	telegram string,
	private int,
) Accounts {
	privateKey, publicKey, err := encrypt.GenRSA()
	if err != nil {
		log.Printf("Failed to generate public and private keys: %v", err)
	}

	id := uuid.New().String()
	return &accounts{
		Uuid:       id,
		Username:   username,
		Password:   password,
		Avatar:     avatar,
		Bio:        bio,
		Name:       name,
		EMail:      EMail,
		Phone:      phone,
		Telegram:   telegram,
		Private:    private,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
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

	acct := &a

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
	acct := &a

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
