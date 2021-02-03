package models

import "github.com/jinzhu/gorm"

type Inbox struct {
	gorm.Model
	Name     	string`gorm:"name"`
	RequestId	string`gorm:"request_id"`
	EventType   string`gorm:"event_type"`
	Actor		string`gorm:"actor"`
}

func NewInboxStructs(i *Inbox) *Inbox {
	r := &Inbox{
		Actor: i.Actor,
		RequestId: i.RequestId,
		EventType: i.EventType,
		Name: i.Name,
	}
	return r
}