package channel

import "gorm.io/gorm"

type channel struct {
	gorm.Model
	name  string   `gorm:"name"`
	owner string   `gorm:"owner"`
	admin []string `gorm:"admin"`
}

type Channel interface {
	New() error
	AddAdmin()
	Update()
}

func newChannel(model gorm.Model, name string, owner string, admin []string) Channel {
	return &channel{Model: model, name: name, owner: owner, admin: admin}
}

func (c *channel) New() error {
	panic("implement me")
	return nil
}

func (c *channel) AddAdmin() {
	panic("implement me")
}

func (c *channel) Update() {
	panic("implement me")
}
