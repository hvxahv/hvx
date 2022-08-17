package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"

	"gorm.io/gorm"
)

const (
	BroadcastsTableName = "broadcasts"
)

type Broadcasts struct {
	gorm.Model

	ChannelId uint `gorm:"primaryKey;type:bigint;channel_id"`

	// ActorId in the AdminId Actor data table, which identifies the creator of the Broadcast.
	// Must be the administrator of the channel.
	AdminId   uint   `gorm:"type:bigint;admin_id"`
	ArticleId uint   `gorm:"type:bigint;article_id"`
	CID       string `gorm:"type:text;cid"`
}

type Broadcast interface {
	Create() error
	GetBroadcasts() ([]*Broadcasts, error) // Get all broadcasts.
	Delete() error
}

func NewBroadcasts(channelId, adminId, articleId uint, cid string) *Broadcasts {
	return &Broadcasts{ChannelId: channelId, AdminId: adminId, ArticleId: articleId, CID: cid}
}

func (b *Broadcasts) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Broadcasts{}); err != nil {
		return errors.NewDatabaseCreate("broadcasts")
	}
	if err := db.Debug().
		Table(BroadcastsTableName).
		Create(&b).
		Error; err != nil {
		return err
	}

	return nil
}

func NewBroadcastsChannelId(channelId uint) *Broadcasts {
	return &Broadcasts{ChannelId: channelId}
}

func (b *Broadcasts) GetBroadcasts() ([]*Broadcasts, error) {
	db := cockroach.GetDB()
	var broadcasts []*Broadcasts

	if err := db.Debug().
		Table(BroadcastsTableName).
		Where("channel_id = ?", b.ChannelId).
		Find(&broadcasts).
		Error; err != nil {
		return nil, err
	}

	return broadcasts, nil
}

func NewBroadcastsDelete(id, channelId, adminId uint) *Broadcasts {
	return &Broadcasts{
		Model:     gorm.Model{ID: id},
		ChannelId: channelId,
		AdminId:   adminId,
	}
}

func (b *Broadcasts) Delete() error {
	isAdmin := NewAdministratesPermission(b.ChannelId, b.AdminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()
	if err := db.Debug().
		Table(BroadcastsTableName).
		Where("id = ?", b.ID).
		Unscoped().
		Delete(&Broadcasts{}).
		Error; err != nil {
		return err
	}

	return nil
}
