package accounts

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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
	Private    int32  `gorm:"private"`
	PrivateKey string `gorm:"type:varchar(3000);private_key"`
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
	// Login Login to the account and generate token, Return token and custom error message.
	Login() (string, error)
}

func NewAccounts(
	username string,
	password string,
	avatar string,
	name string,
	EMail string,
	private int32,
) Accounts {
	// Generate a default public key and private key.
	privateKey, publicKey, err := encrypt.GenRSA()
	if err != nil {
		log.Printf("Failed to generate public and private keys: %v", err)
	}

	// Generate a uuid.
	id := uuid.New().String()

	hash := encrypt.GenPassword(password)

	return &accounts{
		Uuid:       id,
		Username:   username,
		Password:   hash,
		Avatar:     avatar,
		Name:       name,
		EMail:      EMail,
		Private:    private,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

// NewUpdateAccounts Instantiate accounts object.
func NewUpdateAccounts(
	username string,
	password string,
	avatar string,
	bio string,
	name string,
	EMail string,
	phone string,
	telegram string,
	private int32,
) Accounts {
	hash := encrypt.GenPassword(password)

	return &accounts{
		Username:   username,
		Password:   hash,
		Avatar:     avatar,
		Bio:        bio,
		Name:       name,
		EMail:      EMail,
		Phone:      phone,
		Telegram:   telegram,
		Private:    private,
	}
}

// NewAccountByName The accounts object will be instantiated by the username
// for deleting, querying and modifying methods.
func NewAccountByName(username string) *accounts {
	return &accounts{Username: username}
}
func NewAccountLogin(username string, password string) *accounts {
	return &accounts{Username: username, Password: password}
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

func (a *accounts) Login() (string, error) {
	d := db.GetDB()

	var qa *accounts
	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).First(&qa).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(qa.Password), []byte(a.Password)); err != nil {
		return "", errors.Errorf("Password verification failed.")
	}
	token, err := encrypt.GenToken(a.Uuid, a.Username)
	if err != nil {
		log.Println("Token generation failed!")
	}
	return token, nil
}
