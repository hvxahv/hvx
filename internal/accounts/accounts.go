package accounts

import (
	"context"
	"encoding/json"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/disism/hvxahv/pkg/redis"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// AccountData The object tops a user’s profile data and is targeted at GORM.
// Must be a unique key: username, telegram, email, phone.
type AccountData struct {
	gorm.Model
	Uuid       string `gorm:"type:varchar(100);uuid"`
	Username   string `gorm:"primaryKey;type:varchar(100);username;unique"`
	Password   string `gorm:"type:varchar(100);password"`
	Avatar     string `gorm:"type:varchar(100);avatar"`
	Bio        string `gorm:"type:varchar(999);bio"`
	Name       string `gorm:"type:varchar(100);name"`
	EMail      string `gorm:"primaryKey;type:varchar(100);email;unique"`
	Phone      string `gorm:"type:varchar(100);phone;unique"`
	Telegram   string `gorm:"type:varchar(100);telegram;unique"`
	Private    int32  `gorm:"private"`
	PrivateKey string `gorm:"type:varchar(3000);private_key"`
	PublicKey  string `gorm:"type:varchar(3000);public_key"`
}

// Accounts The interface defines the CRUD function for accounts.
type Accounts interface {
	// New Add a user Instantiate using the NewAccounts function.
	New() (int32, error)
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
	// Login Login to the account and generate token, Return token and custom error message.
	Login() (string, error)
}

func NewAccounts(username string, password string, mail string) Accounts {
	// Generate a default public key and private key.
	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("Failed to generate public and private keys: %v", err)
	}

	// Generate a uuid.
	id := uuid.New().String()

	hash := security.GenPassword(password)

	return &AccountData{
		Uuid:       id,
		Username:   username,
		Password:   hash,
		EMail:      mail,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func NewUpdateAcct() *AccountData {
	return &AccountData{}
}

func NewDelAcctByName(name string) Accounts {
	return &AccountData{Username: name}
}

func NewQueryAcctByName(name string) Accounts {
	return &AccountData{Username: name}
}

func NewAccountLogin(name string, password string) Accounts {
	return &AccountData{Username: name, Password: password}
}


func (a *AccountData) New() (int32, error) {
	d := db.GetDB()

	if err := d.AutoMigrate(&AccountData{}); err != nil {
		e := "Automatic table creation failed."
		log.Printf("%s: %v", e, err)
		return 500, errors.Errorf(e)
	}

	acct := &a

	// Before creating, first check whether the user exists. If it does not exist, create the user.
	// If it does, it needs to return an error to the client to explain that the user already exists.
	res := d.Debug().Table("accounts").Where("username = ?", a.Username).First(&acct)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			if err := d.Debug().Table("accounts").Create(&acct).Error; err != nil {
				e := "Failed to create user."
				log.Printf("%s: %v", e, err)
				return 202, errors.Errorf(e)
			}

			// After the user is successfully created,
			// the data encoded by the user's json is stored in the cache,
			// and the cache will never expire.
			ad, _ := json.Marshal(&a)
			_, err := redis.GetRDB().Set(context.Background(), a.Username, ad, 0).Result()
			if err != nil {
				log.Println("Failed to store to cache:", err)
				return 503, errors.New("Failed to store to cache.")
			}

			// Notify the telegram bot that a new user has been added.
			//go func() {
			//	b := bot.NewBot(1, fmt.Sprintf("Added a user: %s", a.Name))
			//	if err := b.Send(); err != nil {
			//		log.Println(err)
			//	}
			//}()

			return 201, errors.Errorf("New account ok!")
		}
	}
	return 202, errors.Errorf("Username already exists.")
}

func (a *AccountData) Query() (*AccountData, error) {
	d := db.GetDB()

	result, err := redis.GetRDB().Get(context.Background(), a.Username).Result()
	if err != nil {
		if err := d.Debug().Table("accounts").Where("username = ?", a.Username).First(&a).Error; err != nil {
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

	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).Updates(&acct).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	}
	return nil
}

func (a *AccountData) Delete() error {
	d := db.GetDB()
	//  Unscoped() Use gorm's Unscoped method to permanently delete data.
	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).Unscoped().Delete(&a).Error; err != nil {
		log.Println(gorm.ErrMissingWhereClause)
		return err
	}
	return nil
}

func (a *AccountData) Login() (string, error) {
	d := db.GetDB()

	var qa *AccountData
	if err := d.Debug().Table("accounts").Where("username = ?", a.Username).First(&qa).Error; err != nil {
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
