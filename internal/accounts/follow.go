package accounts

import (
	"gorm.io/gorm"
)

type Follows struct {
	gorm.Model
	Name       string `gorm:"primaryKey;type:varchar(100);name"`
	TargetName string `gorm:"primaryKey;type:varchar(100);target_name"`
}

func (f *Follows) Followers() {
	panic("implement me")
}

func (f *Follows) Following() {
	panic("implement me")
}

func (f *Follows) New() error {
	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.
	err := NewFollow(f)
	if err != nil {
		return err
	}
	return nil
}

type Follow interface {
	New() error
	Following()
	Followers()
}

func NewFollows(name string, targetName string) Follow {
	return &Follows{Name: name, TargetName: targetName}
}

func NewFoByName(name string) Follow {
	return &Follows{Name: name}
}

func NewFoTargetByName(targetName string) Follow {
	return &Follows{TargetName: targetName}
}
