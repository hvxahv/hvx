package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/APIs/grpc/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

// Accounts is a struct for account.
type Accounts struct {

	// AccountID is a unique identifier for the account.
	gorm.Model

	// Username is the primaryKey of the database which is unique,
	// in the account system of this instance is must be added during
	// the creation process to ensure the correctness of the data.
	Username string `gorm:"primaryKey;type:text;username;unique" validate:"required,min=4,max=16"`

	// Mail When registering, the user is required to provide an email,
	// and an error needs to be returned when the email is not in the
	// correct format. It is also unique in the account system.
	Mail string `gorm:"index;type:text;mail;unique" validate:"required,email"`

	// Password must be encrypted and saved. The length of the password needs to be verified
	Password string `gorm:"type:text;password" validate:"required,min=8,max=100"`

	// ActorID is used for compatibility with the ActivityPub protocol
	// to connect to the actor table by ID.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// IsPrivate sets whether the account is private or not,
	// it is a social extension that is set by the user to make the
	// account public or not.
	IsPrivate bool `gorm:"type:boolean;is_private"`
}

type Account interface {
	IsExist() bool
	Create(publicKey string) error
	Delete() error
	EditUsername(username string) error
	EditPassword(newPassword string) error
	EditEmail(mail string) error
	Verify(password string) (*Accounts, error)
}

// NewUsername ...
func NewUsername(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

// IsExist ...
func (a *Accounts) IsExist() bool {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).Where("username = ? ", a.Username).First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return ok
	}
	return false
}

// GetAccountByUsername ...
func (a *Accounts) GetAccountByUsername() (*Accounts, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(AccountsTable).
		Where("username = ?", a.Username).First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// NewAccounts ...
func NewAccounts(actorID uint, username, mail, password string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: string(hash),
		ActorID:  actorID,
	}
}

// NewAccountsCreate ...
func NewAccountsCreate(username, mail, password string) *Accounts {
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: password,
	}
}

// Create Accounts...
func (a *Accounts) Create(publicKey string) error {
	//if err := validator.New().Struct(a); err != nil {
	//	fmt.Println(err)
	//	return errors.New("FAILED_TO_VALIDATOR")
	//}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		fmt.Println(err)
		return errors.NewDatabaseCreate(serviceName)
	}

	if err := db.Debug().Table(AccountsTable).
		Where("username = ? ", a.Username).Or("mail = ?", a.Mail).
		First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.New(errors.ErrAccountAlready)
		}
	}

	// Create an actor for the account.
	ctx := context.Background()
	client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	if err != nil {
		return err
	}
	create, err := actor.NewActorClient(client.Conn).Create(ctx, &actor.CreateRequest{
		PreferredUsername: a.Username,
		PublicKey:         publicKey,
		ActorType:         "Person",
	})
	if err != nil {
		return err
	}

	actorId, err := strconv.Atoi(create.ActorId)
	if err != nil {
		return err
	}
	v := NewAccounts(uint(actorId), a.Username, a.Mail, a.Password)
	if err := db.Debug().Table(AccountsTable).
		Create(&v).Error; err != nil {
		return fmt.Errorf(errors.ErrAccountCreate)
	}
	return nil
}

// NewAccountsDelete ...
func NewAccountsDelete(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

// Delete ...
func (a *Accounts) Delete() error {
	db := cockroach.GetDB()
	verify, err := NewVerify(a.Username).Verify(a.Password)
	if err != nil {
		return err
	}

	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", verify.ID).
		Unscoped().
		Delete(&Accounts{}).Error; err != nil {
		return err
	}
	return nil
}

func NewAccountsID(id uint) *Accounts {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}

// EditUsername ...
func (a *Accounts) EditUsername(username string) error {
	if ok := NewUsername(username).IsExist(); ok {
		return errors.New("FAILED_TO_VALIDATOR")
	}
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", a.ID).First(&a).Update("username", username).Error; err != nil {
		return err
	}

	address := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), username)
	inbox := fmt.Sprintf("%s/inbox", address)
	if err := db.Debug().Table(ActorsTable).
		Where("id = ?", a.ActorID).
		Update("preferred_username", username).
		Update("inbox", inbox).
		Update("address", address).Error; err != nil {
		return err
	}

	return nil
}

// EditEmail ...
func (a *Accounts) EditEmail(mail string) error {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", a.ID).
		Update("mail", mail).
		Error; err != nil {
		return err
	}

	return nil
}

func NewEditPassword(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

// EditPassword ...
func (a *Accounts) EditPassword(new string) error {
	v, err := NewVerify(a.Username).Verify(new)
	if err != nil {
		return err
	}

	db := cockroach.GetDB()
	hash, _ := bcrypt.GenerateFromPassword([]byte(new), bcrypt.DefaultCost)

	if err := db.Debug().Table(AccountsTable).
		Where("id = ?", v.ID).Update("password", hash).Error; err != nil {
		return err
	}

	return nil
}

func NewVerify(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

func (a *Accounts) Verify(password string) (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table(AccountsTable).Where("username = ?", a.Username).First(&a).Error; err != nil {
		return nil, err
	}

	fmt.Println(password, a.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &Accounts{
		Model: gorm.Model{
			ID: a.ID,
		},
		Username: a.Username,
		Mail:     a.Mail,
		ActorID:  a.ActorID,
	}, nil
}
