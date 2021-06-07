package accounts

import (
	"gorm.io/gorm"
	"hvxahv/pkg/db"
	"log"
)

// accounts ...
type accounts struct {
	gorm.Model
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
	FollowingCount int `gorm:"following_count"`
	FollowerCount  int `gorm:"follower_count"`
}

type Accounts interface {
	New()
	Query()
	Update()
	Delete()
}

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
	privateKey string,
	publicKey string,
	followingCount int,
	followerCount int,
	) Accounts {
	return &accounts{
		Username: username,
		Password: password,
		Avatar: avatar,
		Bio: bio,
		Name: name,
		EMail: EMail,
		Phone: phone,
		Telegram: telegram,
		Social: social,
		Private: private,
		NSFW: NSFW,
		PrivateKey: privateKey,
		PublicKey: publicKey,
		FollowingCount: followingCount,
		FollowerCount: followerCount,
	}
}


func (a *accounts) New() {
	db := db.GetDB()
	if err := db.AutoMigrate(&accounts{}); err != nil {
		return
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


	if err := db.Debug().Table("accounts").Create(&acct).Error; err != nil {
		log.Println(err)
	}

}

func (a *accounts) Query() {
	panic("implement me")
}

func (a *accounts) Update() {
	panic("implement me")
}

func (a *accounts) Delete() {
	panic("implement me")
}
