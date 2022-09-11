package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/APIs/v1alpha1/auth"
	"strconv"

	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/hvxahv/hvx/rsa"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model

	// Username As the primary key of the database, it cannot be duplicated.
	Username string `gorm:"primaryKey;type:text;username;unique" validate:"required,min=4,max=16"`

	// Mail must ensure that the email is unique and properly formatted and cannot be empty.
	// When the user forgets the password, the user is required to provide an email which is unique.
	// Therefore, when registering, the user must provide an email.
	Mail string `gorm:"index;type:text;mail;unique" validate:"required,email"`

	// Password
	// The length of the password must be greater than 8 characters and less than 24 characters, and must not contain spaces.
	Password string `gorm:"type:text;password" validate:"required,min=8,max=24"`

	// When creating a user, an actor needs to be created to record the basic information
	// of the user for compatibility with the ActivityPub protocol.
	// The ActorID field is the id of the actor and is used to associate the actor.
	ActorID uint `gorm:"type:bigint;actor_id"`

	// IsPrivate Set whether the account is private or not, private accounts are not displayed publicly.
	IsPrivate bool `gorm:"type:boolean;is_private"`

	// PrivateKey The private key corresponding to the Activitypub protocol is used
	// only for social activities and no other encryption is involved.
	PrivateKey string `gorm:"type:text;private_key"`
}

type Account interface {
	// IsExist Determine if the account exists.
	IsExist() bool

	// Create need to verify whether the username and email address have been registered or not,
	// and return an error if they have been registered.
	Create(publicKey string) error

	// Delete Verify that the account is correct first, then delete the account by ID.
	Delete() error

	// EditUsername Edit username.
	EditUsername(username string) error

	// EditPassword Edit the password.
	EditPassword(newPassword string) error

	// EditEmail Editorial email.
	EditEmail(mail string) error

	// Verify password is correct.
	Verify(password string) (*Accounts, error)

	// GetPrivateKey When performing activitypub activities,
	// you need to obtain the private key for signing and other operations.
	GetPrivateKey() (string, error)
}

// NewUsername  constructor account's username.
func NewUsername(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

func (a *Accounts) IsExist() bool {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ? ", a.Username).
		First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return ok
	}
	return false
}

func (a *Accounts) GetAccountByUsername() (*Accounts, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ?", a.Username).
		First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

// NewAccounts This constructor is needed to formally create a complete account in the Create Account method.
func NewAccounts(actorID uint, username, mail, password, privateKey string) *Accounts {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &Accounts{
		Username:   username,
		Mail:       mail,
		Password:   string(hash),
		ActorID:    actorID,
		PrivateKey: privateKey,
	}
}

// NewAccountsCreate Constructor for creating an account.
func NewAccountsCreate(username, mail, password string) *Accounts {
	return &Accounts{
		Username: username,
		Mail:     mail,
		Password: password,
	}
}

func (a *Accounts) Create(publicKey string) error {
	// TODO - Verify that the structure data, username email and password match the criteria.
	//if err := validator.New().Struct(a); err != nil {
	//	fmt.Println(err)
	//	return errors.New("FAILED_TO_VALIDATOR")
	//}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Accounts{}); err != nil {
		fmt.Println(err)
		return errors.NewDatabaseCreate(serviceName)
	}

	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ? ", a.Username).
		Or("mail = ?", a.Mail).
		First(&Accounts{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.New(errors.ErrAccountAlready)
		}
	}

	// Create an actor for the account, and return the actor id.
	// Set the type of ActivityPub to Person.
	ctx := context.Background()
	_ := clientv1.New(ctx, microsvc.NewGRPCAddress("actor").Get())
	if err != nil {
		return err
	}
	defer client.Close()

	apk, err := rsa.NewRsa(2048).Generate()
	if err != nil {
		return err
	}
	create, err := actor.NewActorClient(client.Conn).Create(ctx, &actor.CreateRequest{
		PreferredUsername: a.Username,
		PublicKey:         apk.PublicKey,
		ActorType:         "Person",
	})
	if err != nil {
		return err
	}

	actorId, err := strconv.Atoi(create.ActorId)
	if err != nil {
		return err
	}

	v := NewAccounts(uint(actorId), a.Username, a.Mail, a.Password, apk.PrivateKey)
	if err := db.Debug().Table(AccountsTable).
		Create(&v).Error; err != nil {
		return fmt.Errorf(errors.ErrAccountCreate)
	}

	_ := clientv1.New(ctx, microsvc.NewGRPCAddress("auth").Get())
	if err != nil {
		return err
	}
	defer authc.Close()
	spk, err := auth.NewAuthClient(authc.Conn).SetPublicKey(ctx, &auth.SetPublicKeyRequest{
		AccountId: strconv.Itoa(int(v.ID)),
		PublicKey: publicKey,
	})
	if err != nil {
		return err
	}
	if spk.Code != "200" {
		return errors.New(spk.Status)
	}
	return nil
}

// NewAccountsDelete Constructor for deleting an account.
func NewAccountsDelete(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

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

// NewAccountsID Constructing an account ID.
func NewAccountsID(id uint) *Accounts {
	return &Accounts{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (a *Accounts) EditUsername(username string) error {
	if ok := NewUsername(username).IsExist(); !ok {
		return errors.New(errors.ErrAccountUsernameAlreadyExists)
	}
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", a.ID).
		First(&a).
		Update("username", username).Error; err != nil {
		return err
	}

	// Update the actor's preferred username.
	var (
		address = fmt.Sprintf("https://%s/u/%s", viper.GetString("domain"), username)
		inbox   = fmt.Sprintf("%s/inbox", address)
	)

	if err := db.Debug().
		Table(ActorsTable).
		Where("id = ?", a.ActorID).
		Update("preferred_username", username).
		Update("inbox", inbox).
		Update("address", address).Error; err != nil {
		return err
	}

	return nil
}

func (a *Accounts) EditEmail(mail string) error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", a.ID).
		Update("mail", mail).
		Error; err != nil {
		return err
	}

	return nil
}

// NewEditPassword Constructor for editing an account's password.
func NewEditPassword(username, password string) *Accounts {
	return &Accounts{
		Username: username,
		Password: password,
	}
}

func (a *Accounts) EditPassword(new string) error {
	v, err := NewVerify(a.Username).Verify(a.Password)
	if err != nil {
		return err
	}

	db := cockroach.GetDB()
	hash, _ := bcrypt.GenerateFromPassword([]byte(new), bcrypt.DefaultCost)

	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", v.ID).
		Update("password", hash).Error; err != nil {
		return err
	}

	return nil
}

// NewVerify Constructor for verifying an account.
func NewVerify(username string) *Accounts {
	return &Accounts{
		Username: username,
	}
}

func (a *Accounts) Verify(password string) (*Accounts, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AccountsTable).
		Where("username = ?", a.Username).
		First(&a).Error; err != nil {
		return nil, err
	}

	// CompareHashAndPassword to compare whether the password is correct or not.
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

func (a *Accounts) GetPrivateKey() (string, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AccountsTable).
		Where("id = ?", a.ID).
		First(&a).Error; err != nil {
		return "", err
	}

	return a.PrivateKey, nil
}
