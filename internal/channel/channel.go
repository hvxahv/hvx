package channel

import "gorm.io/gorm"

type Channels struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);name"`
	Id        string `gorm:"primaryKey;type:varchar(100);id;unique"`
	Avatar    string `gorm:"type:varchar(999);avatar"`
	Bio       string `gorm:"type:varchar(999);bio"`
	Owner     string `gorm:"primaryKey;type:varchar(100);owner;"`
	Admin     string `gorm:"name"`
	Members   int    `gorm:"type:int;members"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

type Channel interface {
	New() error
	AddAdmin()
	Update()
}

func (c *Channels) New() error {
	panic("implement me")
	return nil
}

func (c *Channels) AddAdmin() {
	panic("implement me")
}

func (c *Channels) Update() {
	panic("implement me")
}
