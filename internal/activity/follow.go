package activity

import "gorm.io/gorm"

type Follows struct {
	gorm.Model

	// Follower actor id.
	Follower uint `gorm:"primaryKey;bigint;follower"`

	// Following actor id.
	Following uint `gorm:"primaryKey;bigint;following"`
}

func (f Follows) Create() error {
	//TODO implement me
	panic("implement me")
}

func (f Follows) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (f Follows) GetFollowers() ([]uint, error) {
	//TODO implement me
	panic("implement me")
}

func (f Follows) GetFollowings() ([]uint, error) {
	//TODO implement me
	panic("implement me")
}

func NewFollows(follower uint, following uint) *Follows {
	return &Follows{Follower: follower, Following: following}
}

type Followee interface {
	Create() error
	Delete() error
	GetFollowers() ([]uint, error)
	GetFollowings() ([]uint, error)
}
